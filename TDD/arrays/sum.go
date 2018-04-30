package main

func sum(n []int) (total int) {
	for _, v := range n {
		total += v
	}
	return
}

func sumAll(args ...[]int) (total []int) {
	for _, arg := range args {
		total = append(total, sum(arg))
	}
	return
}
