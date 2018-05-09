package main

import (
	"fmt"
	"io"
)

func greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
