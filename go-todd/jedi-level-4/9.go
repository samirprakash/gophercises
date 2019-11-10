package main

import "fmt"

func main() {
	m := map[string][]string{}
	m["bond_james"] = []string{"shaken, not stirred", "martini", "women"}
	m["moneypenny_miss"] = []string{"james bond", "literature", "computer science"}
	m["no_dr"] = []string{"being evil", "ice cream", "sun set"}
	m["test_user"] = []string{"test data 1", "test data 2", "test data 3"}

	for k, v := range m {
		fmt.Println(k)
		for i, v := range v {
			fmt.Println("\t", i, v)
		}
	}
}
