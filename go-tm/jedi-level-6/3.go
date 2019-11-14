package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo() has been called")
	fmt.Println("This should have been printed first, but will be printed last due to defer() usage")
	fmt.Println()
}

func bar() {
	fmt.Println("bar() has been called")
	fmt.Println("As per execution, this is called last ans should be printd so but would be printed first as the defer() is in action")
	fmt.Println()
}
