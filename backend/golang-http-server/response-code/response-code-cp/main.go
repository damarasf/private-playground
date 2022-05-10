package main

import (
	"fmt"
	"net/http"
)

var names = []string{
	"Tony",
	"Roger",
	"Bruce",
	"Stephen",
}

func IsNameExists(name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}

	return false
}

func GetNameHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.Method != "GET" {
			http.Error(w, "Method is not allowed.", http.StatusMethodNotAllowed)
			return
		}

		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "Name is required.", http.StatusBadRequest)
			return
		}

		if !IsNameExists(name) {
			http.Error(w, "Name is not exists.", http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "Hello, %s", name)
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/name", GetNameHandler())
	return mux
}
