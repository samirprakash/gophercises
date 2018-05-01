package main

import "fmt"

func main() {
	w := Wallet{}
	w.Deposit(Bitcoin(10))
	fmt.Printf("%s\n", w.Balance())
}
