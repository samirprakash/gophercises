package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(i...))
	fmt.Println(bar(i))
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
