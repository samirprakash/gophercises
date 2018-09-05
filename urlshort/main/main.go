package main

import (
	"fmt"
	"net/http"

	"github.com/samirprakash/gophercises/urlshort"
)

func main() {
	// Fallback URL handler if not path has been provided
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/samirprakash/gophercises/urlshort",
		"yaml-godoc":      "https://godoc.org/gopkg.in/yaml.v2",
	}

	// Process the list of URL's via the mux that has been defined
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Sample YAML inpout data for testing the package implementation
	// yaml := `
	// - path: /urlshort
	// 	url: https://github.com/samirprakash/gophercises/urlshort
	// - path: /urlshort-final
	// 	url: https://github.com/samirprakash/urlshort/gophercises/tree/solution
	// `

	// // Process the YAML input using MApHandler
	// yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

// Default function to handle fallback scenario when no URL has been provided
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

// Hanlde `/` base path from the URL
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
