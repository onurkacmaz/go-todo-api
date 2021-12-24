package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type User struct {
	Name     string
	Email    string
	Password string
}

type Response struct {
	Status int
	Data   []User
}

var users []User
var user User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users = append(users, User{Name: "Onur Kaçmaz", Email: "kacmaz.onur@hotmail.com", Password: "1234"})
	json.NewEncoder(w).Encode(Response{Status: 200, Data: users})
}

func main() {

	host := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	address := net.JoinHostPort(host, port)

	router := mux.NewRouter()
	usersRouter := router.PathPrefix("/api/v1").Subrouter()
	usersRouter.HandleFunc("/users", getUsers).Methods("GET")
	usersRouter.HandleFunc("/users/{id}", getUser).Methods("GET")

	fmt.Printf("server is running at %v", address)
	http.ListenAndServe(address, router)

}
