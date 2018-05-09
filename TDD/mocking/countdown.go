package main

import (
	"fmt"
	"io"
)

func countdown(writer io.Writer) {
	fmt.Fprint(writer, "3\n")
}
