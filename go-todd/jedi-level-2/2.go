package main

import "fmt"

func main() {
	x := 42 == 42
	y := 42 <= 42
	z := 42 >= 42
	i := 42 != 42
	j := 42 < 42
	k := 42 > 42

	fmt.Println(x, y, z, i, j, k)
}
