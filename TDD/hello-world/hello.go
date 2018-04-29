package main

const es = "spanish"
const fr = "french"
const englishSalutation = "Hello, "
const spanishSalutation = "Hola, "
const frenchSalutation = "Bonjour, "

func hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == es {
		return spanishSalutation + name
	}

	if lang == fr {
		return frenchSalutation + name
	}

	return englishSalutation + name
}
