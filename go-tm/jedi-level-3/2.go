package main

import "fmt"

func main() {
	counter := 1
	for i := 65; i <= 90; i++ {
		fmt.Println(counter)
		counter++
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
