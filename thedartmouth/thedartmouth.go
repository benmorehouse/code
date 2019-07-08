package thedartmouth

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

var baseURL = "https://www.thedartmouth.com/"

// thedartmouth News holds all the informaition needed to thedartmouth News scraper
type thedartmouth struct {
	log        *log.Logger
	httpClient *http.Client
	db         *pop.Connection
	fakebox    *fakebox.Client
	retrier    backoff.Retrier
}

// New returns a new thedartmouth scraper
func New(log *log.Logger, httpClient *http.Client, db *pop.Connection, fb *fakebox.Client, retrier backoff.Retrier) *thedartmouth {
	return &thedartmouth{
		log:        log,
		httpClient: httpClient,
		db:         db,
		fakebox:    fb,
		retrier:    retrier,
	}
}

var _ scraper.Scraper = (*thedartmouth)(nil)

// Version returns current version of the scraper
func (b *thedartmouth) Version() string {
	return "thedartmouth, ver. 1.0"
}

//Run executes scraping
func (b *thedartmouth) Run() error {

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
		errors.Wrap(err, "error fetching thedartmouth homepage")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		errors.Wrap(err, "parsing thedartmouth homepage")
	}

	links := []link{}

	// tier 1
	t1selector := doc.Find("article.art-above")
	t1link , _ := t1selector.Find("div.art-above-img-div").Attr("href")
	t1image , err := t1selector.Find("div.art-above-img-div").Attr("src") // this doesnt seem right !
	t1title := t1selector.Find("h6.art-above-headline").Text()
	t1description := t1selector.Find("p.article-abstract").Text()
	links = append(links, link{title: t1title, url: t1url, tier 1, image: t1image, description: t1description})

	//tier 2     Everything is listed the same way after the initial article

	doc.Find("div.row div.col-md-6").Each(func(i int, sel *goquery.Selection) {
		url := sel.Find("article.art-above hed-above-headline").Attr("href")
		title := sel.Find("article.art-above hed-above-headline").Text()
		description := sel.Find("p.article-abstract").Text()
		links = append(links, link{title: title, url: u, tier: 2, description: description})
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
		doc.Find("div.col-md-8 article.main").Each(func(i int, sel *goquery.Selection) {
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
			Source:      "thedartmouth",
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
