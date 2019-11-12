package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println()

	ff := foo()
	fmt.Println(ff())
	fmt.Println(ff())
	fmt.Println(ff())
	fmt.Println(ff())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
