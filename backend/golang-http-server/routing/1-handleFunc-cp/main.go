package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah route untuk TimeHandler dan SayHelloHandler.
// Buatlah route "/time" pada TimeHandler dan "/hello" untuk SayHelloHandler dengan menggunakan http.HandleFunc

var TimeHandler = func(writer http.ResponseWriter, request *http.Request) {
	// TODO: answer here
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	t := time.Now()
	expected := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(expected))
}

var SayHelloHandler = func(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello there")))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}

func main() {
	http.HandleFunc("/time", TimeHandler)
	http.HandleFunc("/hello", SayHelloHandler)
	http.ListenAndServe("localhost:8080", nil)
}
