package middleware

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/controller"
	"strings"
)

type Response struct {
	Status  int
	Message string
}

func ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.Header.Get("Content-Type")
		if c != "application/json" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response{
				Status:  400,
				Message: "Content-Type parameter is required and must be application/json.",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Accept(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.Header.Get("Accept")
		if c != "application/json" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response{
				Status:  400,
				Message: "Accept parameter is required and must be application/json.",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		base64 := strings.Replace(auth, "Basic ", "", -1)
		fmt.Println(base64)
		data, err := b64.StdEncoding.DecodeString(base64)
		if err != nil {
			panic(err)
		}
		credentials := strings.Split(string(data), ":")
		email := credentials[0]
		password := credentials[1]
		if !controller.Check(email, password) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response{
				Status:  401,
				Message: "Invalid credentials",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
