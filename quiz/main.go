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

	correct := 0
	problems := parseRecords(records)
	for i, p := range problems {
		fmt.Printf("Question #%d. %s = ", i+1, p.q)

		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You got %d answers correct out of %d questions\n", correct, len(problems))
}
