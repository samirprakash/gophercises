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

	counter := 1
	cheatingCounter := 1
	low := 1
	high := 100
	guess := 0

	for {
		currentGuess := guess
		guess = (low + high) / 2

		// Check if the user is cheating
		if counter > 1 && currentGuess == guess {
			cheatingCounter++
			if cheatingCounter > 3 {
				fmt.Println()
				fmt.Println("Hey, I think you are cheating! I am going to stop the game now.")
				fmt.Println()
				break
			}
		}

		fmt.Println()
		fmt.Println("I guess the number is:", guess)
		fmt.Println("Is that => ")
		fmt.Println("\t(a) Too high!")
		fmt.Println("\t(b) Too low!")
		fmt.Println("\t(c) Correct!")
		fmt.Println()
		fmt.Print("Your answer => \t")
		scanner.Scan()

		response := scanner.Text()
		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println()
			fmt.Println("I win!")
			fmt.Println()
			break
		} else {
			fmt.Println()
			fmt.Println("This is not a valid input. Please enter: a, b or c!")
			fmt.Println()
		}
		counter++
	}
	fmt.Println("Number of guesses required:", counter)
}
