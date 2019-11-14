package main

import "fmt"

import "encoding/json"

import "os"

type person struct {
	FirstName string
	LastName  string
	Age       int
	Sayings   []string
}

func main() {

	p1 := person{
		FirstName: "person_1_first_name",
		LastName:  "person_1_last_name",
		Age:       21,
		Sayings:   []string{"Saying 1", "Saying 2", "Saying 3"},
	}

	p2 := person{
		FirstName: "person_2_first_name",
		LastName:  "person_2_last_name",
		Age:       22,
		Sayings:   []string{"Saying 4", "Saying 5", "Saying 6"},
	}

	p3 := person{
		FirstName: "person_3_first_name",
		LastName:  "person_3_last_name",
		Age:       23,
		Sayings:   []string{"Saying 7", "Saying 8", "Saying 9"},
	}

	people := []person{p1, p2, p3}

	err := json.NewEncoder(os.Stdout).Encode(&people)
	if err != nil {
		fmt.Println(err)
	}
}
