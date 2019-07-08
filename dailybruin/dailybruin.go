package dailybruin

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gobuffalo/pop"
	"github.com/machinebox/sdk-go/fakebox"
	"github.com/overlooked-incorporated/backend/backoff"
	"github.com/overlooked-incorporated/backend/models"
	"github.com/overlooked-incorporated/backend/scraper"
	"github.com/pkg/errors"
)

var baseURL = "https://www.dailybruin.com/"

// dailybruin News holds all the informaition needed to dailybruin News scraper
type dailybruin struct {
	log        *log.Logger
	httpClient *http.Client
	db         *pop.Connection
	fakebox    *fakebox.Client
	retrier    backoff.Retrier
}

// New returns a new dailybruin scraper
func New(log *log.Logger, httpClient *http.Client, db *pop.Connection, fb *fakebox.Client, retrier backoff.Retrier) *dailybruin {
	return &dailybruin{
		log:        log,
		httpClient: httpClient,
		db:         db,
		fakebox:    fb,
		retrier:    retrier,
	}
}

var _ scraper.Scraper = (*dailybruin)(nil)

// Version returns current version of the scraper
func (b *dailybruin) Version() string {
	return "dailybruin, ver. 1.0"
}

//Run executes scraping
func (b *dailybruin) Run() error {

	if b.httpClient == nil {
		return errors.New("nil HTTP client")
	}

	if b.log == nil {
		return errors.New("nil logger")
	}

	if b.db == nil {
		return errors.New("nil database")
	}

	if b.fakebox == nil {
		return errors.New("nil fakebox")
	}

	if b.retrier == nil {
		return errors.New("nil retrier")
	}

	// home page

	res, err := b.httpClient.Get(baseURL)
	if err != nil {
		errors.Wrap(err, "error fetching dailybruin homepage")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		errors.Wrap(err, "parsing dailybruin homepage")
	}

	links := []link{}

	// tiers are gonna be split into div.large-2,4,6

	// tier 1
	t1selector := doc.Find("div.large-6 div.db-story-c1")
	t1title := t1selector.Find("h2 a").Text()
	t1image := t1selector.Find("div.db-image a").Attr("img")
	t1url := t1selector.Find("div.db-image").Attr("href")
	t1url = baseURL + t1url
	t1description := t1selector.Find("p").Text()

	links = append(links, link{title: t1title, url: t1url, image: t1image, tier: 1, description})

	//tier 2

	doc.Find("div.large-2").Each(func(i int, sel *goquery.Selection) {
		title := sel.Find("div").First().Text()
		u, _ := sel.Find("a.xs-block").First().Attr("href")

		links = append(links, link{title: title, url: u, tier: 2})
	})

	// tier 3

	doc.Find("div.grid-layout-wrapper div.news-feed article").Each(func(i int, sel *goquery.Selection) {
		title := strings.TrimSpace(sel.Find("span.newsblock-story-card__info h2.newsblock-story-card__title").Text())
		u, _ := sel.Find("a.newsblock-story-card__link").First().Attr("href")
		image, _ := sel.Find("img.newsblock-story-card__image ").Attr("src")
		description := sel.Find("p.newsblock-story-card__description").Text()

		links = append(links, link{title: title, url: u, description: description, image: image, tier: 3})
	})

	res.Body.Close()

	var entries []models.Entry

	for i, l := range links {

		b.log.Printf("analyzing %d article out of %d: %s", i+1, len(links), l.url)

		if exists, _ := b.db.Where("url = ?", l.url).Exists(&models.Entry{}); exists {
			continue
		}

		res, err := b.httpClient.Get(l.url)
		if err != nil {
			b.log.Println("error fetching article body:", err.Error())
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			b.log.Println("error parsing article body:", err.Error())
		}

		var content []string
		doc.Find("div.medium-9").Each(func(i int, sel *goquery.Selection) {
			content = append(content, sel.Find("p").Text())
		})

		u, err := url.Parse(l.url)
		if err != nil {
			b.log.Println("error parsing URL")
		}

		text := strings.TrimSpace(strings.Join(content, "\n"))
		if text == "" {
			b.log.Println("no text, continuing...")
			continue
		}

		if len(text) > 3001 {
			text = text[:3000]
		}

		var analysis *fakebox.Analysis
		if err := b.retrier.Do(func() error {
			if l.url != "" {
				analysis, err = b.fakebox.Check(l.title, text, u)
			}
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
			b.log.Println("machinebox fakebox analysis failed:", err.Error())
		}

		description := l.description

		entry := models.Entry{
			Title:       l.title,
			Description: description,
			Content:     text,
			FetchURL:    l.url,
			URL:         l.url,
			Tier:        l.tier,
			Image:       l.image,
			Analyzed:    false,
			Source:      "dailybruin",
		}

		if analysis != nil {
			entry.Analyzed = true
			entry.Score = int(((analysis.Content.Score * 0.7) + (analysis.Title.Score * 0.3)) * 100)
		}

		if err := b.db.Create(&entry); err != nil {
			return errors.Wrap(err, "saving article to the database")
		}

		entries = append(entries, entry)

	}

	b.log.Println("finished scraping articles")

	for _, e := range entries {
		sEntry := scraper.SearchEntry{}
		sEntry.ID = e.ID.String()
		sEntry.Title = e.Title
		sEntry.Content = e.Content

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(&sEntry); err != nil {
			b.log.Println("error marshalling JSON for search:", err.Error())
			continue
		}

		_, err := b.httpClient.Post("http://localhost:8001/index", "application/json", &buf)
		if err != nil {
			b.log.Println("error posting to search:", err.Error())
		}
	}

	b.log.Println("articles indexed for search")

	return nil

}

type link struct {
	title       string
	url         string
	tier        int
	image       string
	description string
}
