package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Membaca cookie token kemudian return cookie kedalam body response
	mux.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		cookieFieldName := "token"

		// Task:  1. Ambil token dari cookie yang dikirim ketika request
		// 		  2. return unauthorized ketika token kosong
		// 		  3. return bad request ketika field token tidak ada

		req := r.Header.Get(cookieFieldName)
		for _, cookie := range r.Cookies() {
			if cookie.Name == cookieFieldName {
				req = cookie.Value
			}
		}

		if req == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		if req == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}

		w.Write([]byte(fmt.Sprintf("Tokenmu adalah %s!", req)))

		// req := r.Cookies()
		// for _, cookie := range req {
		// 	if cookie.Name == cookieFieldName {
		// 		fmt.Fprintf(w, "Tokenmu adalah %s!", cookie.Value)
		// 		return
		// 	}
		// }
		// w.WriteHeader(http.StatusUnauthorized)
		// fmt.Fprintf(w, "Unauthorized")
	})

	return mux
}
