package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, cobalah untuk melakukan chaining middleware.
// Lengkapi middleware AdminMiddleware dan RequestMethodGetMiddleware.
// Pada AdminMiddleware, lakukan pengecekan terhadap header "role". Jika nilai header "role" tidak sama dengan "ADMIN", maka return error HTTP Unauthorized.
// Pada RequestMethodGetMiddleware, lakukan pengecekan terhadap request method. Jika request method selain GET, maka return error HTTP Status Not Allowed.

func TimeHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		output := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
		fmt.Fprint(writer, output)
	}
}

func RequestMethodGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			log.Println("Method not allowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
	}) // TODO: replace this
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Role") != "ADMIN" {
			log.Println("Unauthorized")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}) // TODO: replace this
}

func main() {
	// TODO: answer here
	handler := TimeHandler()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: AdminMiddleware(RequestMethodGetMiddleware(handler)),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
