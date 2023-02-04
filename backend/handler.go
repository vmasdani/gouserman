package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router) {

	// Routes consist of a path and a handler function.
	r.HandleFunc("/hello", YourHandler)

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		secret, err := os.ReadFile("jwt.secret")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "secret not found")
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"jti": "bar",
			"nbf": time.Unix(time.Now().Unix()+86400*365, 0),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(secret)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Failed signing token")
		}

		fmt.Fprint(w, tokenString)
	}).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))
}
