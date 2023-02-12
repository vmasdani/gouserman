package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/samber/lo"
	"google.golang.org/protobuf/encoding/protojson"
)

func Route(r *mux.Router) {

	// Routes consist of a path and a handler function.
	r.HandleFunc("/hello", YourHandler).Methods(http.MethodGet)

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
	}).Methods(http.MethodPost)

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var user GousermanUser
		protojson.Unmarshal(b, &user)

		DB.Save(&User{
			Username: lo.ToPtr("Testuser" + time.Now().Format(time.RFC3339)),
			Email:    lo.ToPtr("Testemail" + time.Now().Format(time.RFC3339)),
		})

		s, _ := protojson.Marshal(&user)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(s))
	}).Methods(http.MethodPost)

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		u := []User{}
		DB.Find(&u)

		users := GousermanUsers{
			Users: lo.Map(u, func(u User, i int) *GousermanUser {
				return lo.ToPtr(GousermanUser{
					GormModel: &GousermanGormModel{
						Id: lo.ToPtr(uint64(u.ID)),
					},
					Username: u.Username,
					Password: u.Password,
					Email:    u.Email,
				})
			}),
		}

		s, _ := protojson.Marshal(&users)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(s))
	}).Methods(http.MethodGet)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))
}
