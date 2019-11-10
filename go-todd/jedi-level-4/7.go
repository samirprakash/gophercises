package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, Not Stirred"}
	mp := []string{"Miss", "MoneyPenny", "Hello! James"}
	x := [][]string{jb, mp}

	for _, v := range x {
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}
