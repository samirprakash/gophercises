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

	guess := 50
	for {
		fmt.Println("I guess the number is: ", guess)
		fmt.Println("Is that => ")
		fmt.Println("(a) Too high!")
		fmt.Println("(b) Too low!")
		fmt.Println("(c) Correct!")
		scanner.Scan()

		response := scanner.Text()
		if response == "a" {
			guess--
		} else if response == "b" {
			guess++
		} else if response == "c" {
			fmt.Println("I win!")
			break
		} else {
			fmt.Println("This is not a valid input. Please try again!")
		}
	}
}
