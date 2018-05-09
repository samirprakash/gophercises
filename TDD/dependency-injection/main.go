package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Testing Dependency Injection!")
	greet(os.Stdout, "Samir")

	http.ListenAndServe(":5000", http.HandlerFunc(greetHandler))
}
