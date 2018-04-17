package main

import (
	"encoding/csv"
	"flag"
	"fmt"
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
