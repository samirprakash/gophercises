package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       uint8
}

func changeMe(p *person) {
	(*p).firstName = "aashi"
	(*p).lastName = "bhushan"
	p.age = 30
}

func main() {
	p1 := person{
		firstName: "samir",
		lastName:  "prakash",
		age:       35,
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
