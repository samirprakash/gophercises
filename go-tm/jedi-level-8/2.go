package main

import (
	"encoding/json"
	"fmt"
)

// Person defines the attributes for a person
type Person struct {
	FirstName string   `json:"FirstName"`
	LastName  string   `json:"LastName"`
	Age       int      `json:"Age"`
	Sayings   []string `json:"Sayings"`
}

func main() {
	s := `[{"FirstName": "Samir", "LastName": "Prakash", "Age": 35, "Sayings": ["Saying 1", "Saying 2", "Saying 3"]}, {"FirstName": "Aayushi", "LastName": "Bhushan", "Age": 30, "Sayings": ["Saying 4", "Saying 5", "Saying 6"]}]`
	fmt.Println(s)

	people := []Person{}
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
