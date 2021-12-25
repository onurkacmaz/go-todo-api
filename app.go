package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"os"
	"time"
)

type User struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func getUsers(db *sql.DB) http.HandlerFunc {

	type Response struct {
		Status int
		Users  []User
	}

	fn := func(w http.ResponseWriter, r *http.Request) {

		var users []User

		rows, err := db.Query(`SELECT * FROM users ORDER BY created_at DESC`)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var u User
			err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
			if err != nil {
				panic(err)
			}
			users = append(users, u)
		}
		err = rows.Err()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Status: 200,
			Users:  users,
		})
	}

	return fn

}

func getUser(db *sql.DB) http.HandlerFunc {

	type Response struct {
		Status int
		User   User
	}

	fn := func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		requestedId := params["id"]

		var (
			id        int
			name      string
			email     string
			password  string
			createdAt time.Time
		)

		err := db.QueryRow(`SELECT * FROM users WHERE id = ?`, requestedId).Scan(&id, &name, &email, &password, &createdAt)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Status: 200,
			User: User{
				Id:        requestedId,
				Name:      name,
				Email:     email,
				Password:  password,
				CreatedAt: createdAt,
			},
		})
	}
	return fn
}

func main() {

	host := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	address := net.JoinHostPort(host, port)

	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/go_rest?parseTime=true")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	usersRouter := router.PathPrefix("/api/v1").Subrouter()
	usersRouter.HandleFunc("/users", getUsers(db)).Methods("GET")
	usersRouter.HandleFunc("/users/{id}", getUser(db)).Methods("GET")

	fmt.Printf("server is running at %v \n", address)
	http.ListenAndServe(address, router)

}
