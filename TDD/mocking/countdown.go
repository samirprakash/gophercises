package main

import (
	"bytes"
	"fmt"
)

func countdown(writer *bytes.Buffer) {
	fmt.Fprintf(writer, "3")
}
