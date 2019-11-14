package main

import "fmt"

func main() {
	s := struct {
		firstName string
		lastName  string
		addresses map[string]string
		aliases   []string
	}{
		firstName: "Samir",
		lastName:  "Prakash",
		addresses: map[string]string{"home": "address 1", "work": "address 2"},
		aliases:   []string{"sonu", "chamanchindi"},
	}

	fmt.Println(s.firstName, s.lastName)
	for k, v := range s.addresses {
		fmt.Println(k, "-", v)
	}
	for i, v := range s.aliases {
		fmt.Println("Alias", i, v)
	}
}
