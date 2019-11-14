package main

import "fmt"

const (
	currentYear = 2019
	_           = iota
	a           = currentYear - iota
	b           = currentYear - iota
	c           = currentYear - iota
	d           = currentYear - iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
