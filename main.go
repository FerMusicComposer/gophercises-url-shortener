package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	// build the mapHandler using mux as fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := mapHandler(pathsToUrls, mux)

	// build yamlHandler using mapHandler as fallback

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi there!")
}
