package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSum(sum, xi)
}

func printSum(f func([]int) int, xi []int) {
	fmt.Println(f(xi))
}

func sum(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
