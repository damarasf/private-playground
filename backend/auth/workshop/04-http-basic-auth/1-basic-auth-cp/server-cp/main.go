package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var (
	username = "aditira"
	password = "aditira123"
)

func main() {
	fmt.Println("Starting Server at port :8080")
	http.ListenAndServe(":8080", Routes())
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			// TODO: answer here
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Error parsing basic auth")
			return
		}
		if u != username {
			// TODO: answer here
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "{\"message\": \"Invalid username\"}")
			return
		}
		if p != password {
			// TODO: answer here
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "{\"message\": \"Invalid password\"}")
			return
		}
		fmt.Printf("Username: %s\n", u)
		fmt.Printf("Password: %s\n", p)

		// TODO: answer here
		token, err := bcrypt.GenerateFromPassword([]byte(u), bcrypt.MinCost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \"Error encrypting token\"}")
			return
		}

		cookie := http.Cookie{
			Name:  "token",
			Value: string(token),
			Path:  "/",
		}
		http.SetCookie(w, &cookie)
		fmt.Fprint(w, "{\"message\": \"welcome to CAMP 2022!\"}")
	})

	return mux
}

// Encode auth aditira:aditira123 disini -> https://www.base64encode.org/

// $ curl -v -X POST http://localhost:8080/login -H "Authorization: Basic YWRpdGlyYTphZGl0aXJhMTIz"
