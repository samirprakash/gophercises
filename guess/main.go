package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Guess a number between 1 and 100")
	fmt.Println("Press ENTER when ready!")
	scanner.Scan()

	counter := 0
	low := 1
	high := 100

	for {
		guess := (low + high) / 2
		counter++
		fmt.Println("I guess the number is:", guess)
		fmt.Println("Is that => ")
		fmt.Println("(a) Too high!")
		fmt.Println("(b) Too low!")
		fmt.Println("(c) Correct!")
		scanner.Scan()

		response := scanner.Text()
		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I win!")
			break
		} else {
			fmt.Println("This is not a valid input. Please try again!")
		}
	}
	fmt.Println("Number of guesses required:", counter)
}
