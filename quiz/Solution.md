
We should try to break the problem into smaller parts so that we can handle these requirements with proper execution and test.

* Requirement 1
  * Input to the quiz program needs to be read from a CSV file.
  * Print content of the CSV file on the console.
* Requirement 2
  * Iterate through the content of the files and generate questions and answers for the user.
  * Provide questions to the user and an option to submit answers.
  * Evaluate these answers and provide a result to the user.
* Requirement 3
  * Add a timer flag with a default value of 30 seconds
	* Use the timer to exit from the program once the time limit has expired

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

```go
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
1. Parse the `[][]string` records from CSV into a `[]problem` using the `parseRecords()` present in the `parse.go` file. 
1. `problem` struct has been defined at the top with two fields - `q for question and a for answer`. `parseRecords()` would generate a slice of type `problem`.
1. Iterate over `[]problem` and print one question at a time on the console
1. Using `fmt.Scanf()`, wait for the user to provide an answer and then capture that answer
1. Compare that answer from the CSV file and increment the `correct` counter for each correct answer
1. Print the result at the end of the loop after exhausting all the questions.

```go
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
	...
	...

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
---
### Requirement 3

For adding a timer to the existing program, we would be using the `time` package which provides a handy `timer` type. There is a function `NewTimer` in this type, which returns a new timer after a fixed duration on a channel. We would be using this channel to check if we need to exit from the program.

Another issue that we will address here is the execution pause that happens due to using `fmt.Scanf()`. This pauses the program till the time the user provides as input. Unless the user provides an input, the program execution will not proceed. This does not go well with our auto-exit logic when the time limit expires. We would be handling this issue by using go routines.

The code snippet added below just shows the addition we have done over `requirement 2`:
1. We have added a new variable `timer` which returns a new timer after the user defined time `limit` has expired
1. We use this return value channel `<-timer.C` to select when the program exits
1. For handling the `fmt.Scanf()` pause problem, we have removed that logic into a go routing anonymous function which returns the user inout into `ach`, another channel.
1. We then use this logic to selec the condition where if a value is received in the `<-ach`, we would be adding that to the list of answers received.
1. However, by executing the user input logic as a go routine, this flow executes in parallel to the timer logic in the program
1. Based on which input is received, the program make a decision to continue with more questions or to exit

```go
package main

import (
	...
	...
	"time"
)

func main() {
	limit := flag.Int("limit", 30, "maximum time to answer a question before user is time out of the game")
	flag.Parse()

	...
	...
	...

	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	correct := 0
	problems := parseRecords(records)

	for i, p := range problems {
		fmt.Printf("Question #%d. %s = ", i+1, p.q)

		ach := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ach <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou got %d answers correct out of %d questions\n", correct, len(problems))
			return
		case answer := <-ach:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You got %d answers correct out of %d questions\n", correct, len(problems))
}
```

# Further Reading

* [package string](https://golang.org/pkg/strings/)
* [package time](https://golang.org/pkg/time/)
* [package os](https://golang.org/pkg/os/)
* [package fmt](https://golang.org/pkg/fmt/)
* [package encoding/csv](https://golang.org/pkg/encoding/csv/)
* [Go Channels](https://golang.org/doc/effective_go.html#channels)
* [Go Routines](https://golang.org/doc/effective_go.html#goroutines)
* [Original problem statement](https://gophercises.com/exercises/quiz)