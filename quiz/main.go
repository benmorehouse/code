package main

import(
	"flag"
	"strings"
	"fmt"
	"log"
	"os"
	"encoding/csv"
	"time"
)

func main(){
	csv_filename := flag.String("csv","problems.csv", "a csv file with our data") // this is a pointer to problems.csv
	flag.Parse()
	// this also makes the first argument available for use to use in the rest of the code
	file , err := os.Open(*csv_filename)
	if err != nil{
		log.Println("Failed to open ", *csv_filename)
	}
	// we are gonna make a new reader using this csv file and then loop through it and keep track of how many times the user is correct
	r := csv.NewReader(file) // csv takes in an ioreader and returns a new reader for csv files
	lines , err := r.ReadAll() // this will then take the csv file reader and read through and return a string of all the objects that we are trying to touch
	if err != nil{
		log.Println(err)
	}
	problems := parse_lines(lines) // problems is of type problem set

	correct := 0
	var input string

	for _ , val := range problems{
		fmt.Println(val.question)
		fmt.Scan(&input)
		if input == (val.answer){
			correct++
		}
	}
	fmt.Println("correct is",correct)
}

func parse_lines(lines [][]string) []problem_set{
	returnVal := make([]problem_set,len(lines))
	for i, val := range lines{
		returnVal[i] = problem_set{
			question: val[0],
			answer: strings.TrimSpace(val[1]),
		}
	}
	return returnVal
}

type problem_set struct{
	question string
	answer string
}

func timer(input int)chan int{
	time.Sleep(input * time.Seconds()
}


