/*Call back function example*/
package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	si := sum(ii...)
	sei := sumOfEven(sum, ii...)
	soi := sumOfOdd(sum, ii...)
	fmt.Println(si)
	fmt.Println(sei)
	fmt.Println(soi)

}

func sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func sumOfEven(f func(x ...int) int, i ...int) int {
	yi := []int{}
	for _, v := range i {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func sumOfOdd(f func(x ...int) int, i ...int) int {
	yi := []int{}
	for _, v := range i {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
