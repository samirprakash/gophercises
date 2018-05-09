package main

import (
	"fmt"
	"io"
	"net/http"
)

func greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	greet(w, "World")
}
