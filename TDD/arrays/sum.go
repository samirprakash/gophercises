package main

func sum(n [5]int) (sum int) {
	for _, v := range n {
		sum += v
	}
	return
}
