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

	switch lang {
	case es:
		return spanishSalutation + name
	case fr:
		return frenchSalutation + name
	default:
		return englishSalutation + name
	}
}
