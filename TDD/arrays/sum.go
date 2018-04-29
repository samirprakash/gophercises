package main

func sum(n [5]int) (total int) {
	for i := 0; i < 5; i++ {
		total += n[i]
	}
	return
}
