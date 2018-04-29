package main

const pe = "Hello, "
const ps = "Hola, "

func hello(s, l string) string {
	if s == "" {
		s = "World"
	}

	if l == "spanish" {
		return ps + s
	}

	return pe + s
}
