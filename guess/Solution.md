# Solution

The problem at hand has got 4 distinct section that can be handled one at a time to arrive at the final solution.

- Create a program to guess a number from 1 to 100 proceeding linearly based on user input
- Optimize the program with a search algorithm so that number of guesses are reduced
- Find out the number of guesses required to guess the correct number that the user has selected
- Check if the user is cheating and exit the program if that is the case

### Create a program to guess a number from 1 to 100 proceeding linearly based on user input

Program needs to make an initial guess between 1 and 100 to get the execution flow going. It also needs to be provide user with three options - high, low or correct - in order to decide what the next guess would be based on the user's input.

`bufio` and `os` packages can be used to get the user input. The guess can be based on adding 1 to the guess if the user says that the previous guess made by the program was `too low` or removing 1 to the guess if the user chooses otherwise.

Below program does that. We provide three options to the user and we set the original guess to `50`, a number mid-way 1 and 100. We use `scanner` to get the input from user and then use that input to increment or decrement the guess by 1.

We exit the program when the user chooses the option `(c)` for the guess being correct.

```go
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

	for {
		guess = 50

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
			guess = guess - 1
		} else if response == "b" {
			guess = guess + 1
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
	}
}
```

### Optimize the program with a search algorithm so that number of guesses are reduced
### Find out the number of guesses required to guess the correct number that the user has selected

If you execute the above program, it would work as it is expected to. However, it will take a long time to guess the correct answer due the increment/decrement with 1 for every option the user chooses. We need to optimize this by using a search algorithm that reduces the complexity of this program. For that, we would be using the binary search algorithm.

All we need to do is define `low` and `high` as two variables which will be the lower and upper bound of search range. Then we set the initial guess to `(low + high) \ 2` and based on the choice selected by the user, we reset `high = guess - 1` and `low = guess + 1`. 

Every iteration handles the upper and lower bound based on the user choice, which in turn would update the next guess by calculating it based on the new values for `low` and `high`.

For finding out the number of guesses required, we would just add a counter initialized to 0 and then increment the counter after every successful pass of the `for` loop.

Below program shows the modifications on the previous program.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  ...
  ...

  counter := 0
	low := 1
	high := 100
	guess := 0

	for {
		guess = (low + high) / 2

      ...
      ...

		response := scanner.Text()
		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
        ...
        ...
    }
    counter++
	}
	fmt.Println("Number of guesses required:", counter)
}
```

### Check if the user is cheating and exit the program if that is the case

Once we have a fully functional number guessing program, one of the things that we want to check for is the condition when the user is cheating with the program.

By cheating, I mean that say the user selected the number 70 between 1 and 100. But instead of replying truthfully to the provided choices, the user intentionally provides wrong choice. This would mislead the program because the next guess that the program generates is based on the truthful answer to the questions - low, high or correct - by the user. We need to handle this and exit the program if the user continues to lie.

The way we can determine the logic for this lies in the implementation of the binary search algorithm. Since we have reduced the number of iterations required by the program to arrive to the correct answer by using the algorithm, if the user keeps on cheating and providing incorrect choices then very soon we can see that the guess from the program will reach a number that will be same for every next execution.

We can use this to pin the user and tell him that cheating is now allowed. We exit the program after that.

Below is the logic for this check.

We define a `cheatingCounter` and `currentGuess` and assign them the values 1 and of `guess` respectively. We check that at least one iteration of the program has executed and then we check if the `currentGuess == guess`. If this condition is met then we increment the cheating counter. When the cheating counter reaches 3, then we inform the user and exit.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  ...
  ...

	counter := 1
	cheatingCounter := 1
	
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

    ...
    ...
}
```