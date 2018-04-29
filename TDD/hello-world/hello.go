package main

const p = "Hello, "

func hello(s string) string {
	if s == "" {
		s = "World"
	}
	return p + s
}
