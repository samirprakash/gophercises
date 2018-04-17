
We should try to break the problem into smaller parts so that we can handle these requirements with proper execution and test.

* Requirement 1
  * __Input to the quiz program needs to be read from a CSV file.__
  * __Print content of the CSV file on the console.__
* Requirement 2
  * __Iterate through the content of the files and generate questions and answers for the user__
  * __Provide questions to the user and an option to submit answers__
  * __Evaluate these answers and provide a result to the user__
* Requirement 3
  * __Add a timer with a default setting of 30 seconds__

---
### Requirement 1

By default, we have created a `problems.csv` file which contains multiple rows of questions and answers. But, a quiz game should be able to take different input files and still be able to run. We would like to provide the quiz master with an option to provide different Q&A inputs to our program. 

This can be achieved using the `flag` package. This package allows us to configure and pass values to the executable in order to modify its execution - based on the flags and their values.

If you have a `problems.csv` in your current directory and you create a `main.go` with the below mentioned content, executing `go install && quiz -h` will display `-csv` as a flag for your executable having a default value of `problems.csv`.

Execute `go install && quiz -csv=problems.csv` to check the program in action.

Going through the code below: 
1. Read a user defined string input into a flag named `csv` with a default value of `problems.csv`. 
1. Use the `flags.Parse()` function to make this flag available to our executable.
1. Try to open this file with `os.Open()` and handle error, if any
1. If file is opened successfully, then read its content and print it on the console using the `encoding/csv` package's `ReadAll()` function.

```
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
  
  file, err := os.Open(*csvFilename)
	if err != nil {
		// handle err
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		//handle err
	}
  fmt.Println(records)
}
```
---
### Requirement 2

After reading the content of the CSV file, we need to parse the entries into questions and answers and provide them to the end user in an iterative manner so that they can analyze the questions and provide answers to your program for evaluation.

For this we need to convert `[][]string` being returned by `ReadAll()` into a slice `[]problem`, This `problem` struct contains two fields - question and answer. Once we have converted the content of `[][]string` into `[]problem`, we then need to iterate over `[]problem` and display the questions to the end user in an orderly fashion.

Once a question has been displayed, user is going to provide an answer before we proceed to the next question. We need to capture this answer using `fmt.Scanf()`. Afyter capturing this value, we would check with the existing list of answers and validate if the user has prvovided a correct or a wrong answer. We would also keep a counter for the number of right answers and the end of the quiz, will display the tally on the console for the user.

Modifying the existing program as such would be providing the desired execution:
1. Declare a counter `correct`
1. Parse the `[][]string` records from CSV into a `[]problem`. `problem` struct has been defined at the top with two fields - `q for question and a for answer`.
1. Iterate over `[]problem` and print one question at a time on the console
1. Using `fmt.Scanf()`, wait for the user to provide an answer and then capture that answer
1. Compare that answer from the CSV file and increment the `correct` counter for each correct answer
1. Print the result at the end of the loop after exhausting all the questions.

```
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
```

### Requirement 3