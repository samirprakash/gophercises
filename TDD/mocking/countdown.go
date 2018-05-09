package main

import (
	"fmt"
	"io"
)

const (
	startCounter int8   = 3
	exitCode     string = "Go!"
)

func countdown(writer io.Writer) {
	for i := startCounter; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, exitCode)
}
