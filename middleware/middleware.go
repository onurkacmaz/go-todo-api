package middleware

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int
	Message string
}

func ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.Header.Get("Content-Type")
		if c != "application/json" {
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
			json.NewEncoder(w).Encode(Response{
				Status:  400,
				Message: "Accept parameter is required and must be application/json.",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
