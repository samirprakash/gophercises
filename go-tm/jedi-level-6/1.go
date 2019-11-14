package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	y, s := bar()
	fmt.Println(y)
	fmt.Println(s)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "function called"
}
