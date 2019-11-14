package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	flavors   []string
}

func main() {
	p1 := person{
		firstName: "Samir",
		lastName:  "Prakash",
		flavors:   []string{"vanilla", "rum and coke"},
	}

	p2 := person{
		firstName: "Aayushi",
		lastName:  "Bhushan",
		flavors:   []string{"vanilla", "chocolate", "Strawberry"},
	}

	m := map[string]person{}
	m[p1.lastName] = p1
	m[p2.lastName] = p2

	for _, v := range m {
		fmt.Println(v.firstName, v.lastName)
		for _, vv := range v.flavors {
			fmt.Println("\t", vv)
		}
	}
}
