package main

import (
	"fmt"
	"net/http"

	"github.com/samirprakash/gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/url-shortner-godoc": "https://godoc.org/github.com/samirprakash/urlshort",
		"yaml-godoc":          "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
	- path: /urlshort
		url: https://github.com/samirprakash/urlshort
	- path: /urlshort-final
		url: https://github.com/samirprakash/urlshort/tree/solution
	`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World")
}
