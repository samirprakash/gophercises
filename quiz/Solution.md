
We should try to break the problem in smaller parts so that we can handle these requirements with proper execution and test.

* Requirement 1
  1. __Input to the quiz program needs to be read from a CSV file.__
  1. __Print content of the CSV file on the console.__
* Requirement 2
  1. __Iterate through the content of the files and generate questions and answers for the user__
  1. __Provide questions to the user and an option to submit answers__
  1. __Evaluate these answers and provide a result to the user__
* Requirement 3
  1. __Add a timer with a default setting of 30 seconds__


### Requirement 1

By default we have created a `problems.csv` file which contains multiple rows of questions and answers. But a quiz game should be able to take different input files and still be able to run.We would like to provide the quiz master with an option to provide different Q&A inputs to our program. 

This can be achieved using the `flag` package. This package allws us to configure and pass values to the executable probabaly in order to modify its execution based on the flags and their values.

If you have a `problems.csv` in your current directoryh and you create a `main.go` with the below mentioned content, executing `go install && quiz -h` will display `-csv` as a flag for your executable having a default value of `problems.csv`.

Going through the code below: 
1. Read a user defined string input into a flag named `csv` with a default value of `problems.csv`. 
1. Use the `flags.Parse()` function to make this flag available to our executable.
1. Try to open this file with `os.Open()` and handle error, if any
1. If file is opened successfully, then read its content and print it on the console.

> r.ReadAll() returns a [][]string which needs to be handled.

```
package main

import (
	"encoding/csv"
	"flag"
  "os"
  "fmt"
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

### Requirement 2