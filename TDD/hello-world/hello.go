package main

const spanish = "spanish"
const englishSalutation = "Hello, "
const spanishSalutation = "Hola, "

func hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == spanish {
		return spanishSalutation + name
	}

	return englishSalutation + name
}
