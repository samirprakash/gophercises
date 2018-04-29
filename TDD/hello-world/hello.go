package main

func hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return salutation(lang) + name
}
