package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Testing Dependency Injection!")
	greet(os.Stdout, "Samir")
}
