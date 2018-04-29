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
	return salutation(lang) + name
}

func salutation(lang string) (salutation string) {
	switch lang {
	case es:
		salutation = spanishSalutation
	case fr:
		salutation = frenchSalutation
	default:
		salutation = englishSalutation
	}
	return
}
