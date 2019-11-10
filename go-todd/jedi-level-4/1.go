package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for _, v := range a {
		fmt.Printf("%T\t%v\n", v, v)
	}
	fmt.Printf("%T\n", a)
}
