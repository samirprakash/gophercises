package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) speak() {
	fmt.Println("I am", p.firstName, p.lastName, "and my age is", p.age)
}

func main() {
	p := person{
		firstName: "Samir",
		lastName:  "Prakash",
		age:       35,
	}
	p.speak()
}
