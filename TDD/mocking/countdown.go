package main

import (
	"fmt"
	"io"
	"time"
)

const (
	startCounter int8   = 3
	exitCode     string = "Go!"
)

func countdown(writer io.Writer) {
	for i := startCounter; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(writer, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(writer, exitCode)
}
