package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("not able to open the file %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("not able to open the file %s\n", err))
	}

	problems := parseRecords(records)
	fmt.Println(problems)
}

func parseRecords(records [][]string) []problem {
	ret := make([]problem, len(records))
	for i, r := range records {
		ret[i] = problem{
			q: r[0],
			a: r[1],
		}
	}
	return ret
}

func exit(m string) {
	log.Println(m)
	os.Exit(1)
}
