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
		flavors:   []string{"vanilla", "falooda"},
	}

	p2 := person{
		firstName: "Aayushi",
		lastName:  "Bhushan",
		flavors:   []string{"vanilla", "chocolate"},
	}

	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.flavors {
		fmt.Println("\t", v)
	}
	fmt.Println(p2.firstName, p2.lastName)
	for _, v := range p2.flavors {
		fmt.Println("\t", v)
	}
}
