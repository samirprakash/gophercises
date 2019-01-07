package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	// define the flags - csv and limit
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := flag.Int("limit", 30, "maximum time to answer a question before user is timed out of the game")
	flag.Parse()

	// open and read CSV
	file, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatalf("not able to open the file %s\n", *csvFilename)
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("not able to open the file %s\n", err)
	}

	// generate a new timer based on limit provided by the user
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	correct := 0
	// convert [][]string to []problem
	problems := parseRecords(records)

	for i, p := range problems {
		fmt.Printf("Question #%d. %s = ", i+1, p.q)

		// initialize a go routine and get the result back in a channel - ach
		ach := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ach <- answer
		}()

		// select condition based on whether time expires or the user provides an input within the stipulated time
		select {
		case <-timer.C:
			fmt.Printf("You got %d answers correct out of %d questions\n", correct, len(problems))
			return
		case answer := <-ach:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("You got %d answers correct out of %d questions\n", correct, len(problems))
}
