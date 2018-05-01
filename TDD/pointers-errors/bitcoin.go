package main

import "fmt"

// Bitcoin is the type present in wallet
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
