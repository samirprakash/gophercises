package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Testing")
	cmd := exec.Command("echo", "hello")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
