package main

import "fmt"

func main() {
	x := r()
	fmt.Println(x())
}

func r() func() int {
	return func() int {
		return 1
	}
}
