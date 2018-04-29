package main

const es = "spanish"
const fr = "french"
const defaultSalutation = "Hello, "
const spanishSalutation = "Hola, "
const frenchSalutation = "Bonjour, "

func salutation(lang string) (salutation string) {
	switch lang {
	case es:
		salutation = spanishSalutation
	case fr:
		salutation = frenchSalutation
	default:
		salutation = defaultSalutation
	}
	return
}
