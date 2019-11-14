package main

import "fmt"

import "encoding/json"

// User defines the attributes of an user for the system
type User struct {
	FirstName string
	Age       int
}

func main() {
	u1 := User{
		FirstName: "samir",
		Age:       35,
	}

	u2 := User{
		FirstName: "aashi",
		Age:       30,
	}

	users := []User{u1, u2}
	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))
}
