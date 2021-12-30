package middleware

import (
	"encoding/json"
	"net/http"
	"rest-api/controller"
	"rest-api/util"
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

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
		if !controller.IsTokenValid(token) {
			util.Response{
				Status:  401,
				Message: "Token is invalid.",
			}.ResponseJson(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
