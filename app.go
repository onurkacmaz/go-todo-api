package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt string
}

type Response struct {
	Status int
	User   User
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
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				panic(err)
			}
		}(rows)

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
		err = json.NewEncoder(w).Encode(Response{
			Status: 200,
			Users:  users,
		})
		if err != nil {
			return
		}
	}

	return fn

}

func showUser(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		requestedId := params["id"]

		var (
			id        int
			name      string
			email     string
			password  string
			createdAt string
		)

		err := db.QueryRow(`SELECT * FROM users WHERE id = ?`, requestedId).Scan(&id, &name, &email, &password, &createdAt)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(Response{
			Status: 200,
			User: User{
				Id:        id,
				Name:      name,
				Email:     email,
				Password:  password,
				CreatedAt: createdAt,
			},
		})
		if err != nil {
			return
		}
	}
	return fn
}

func deleteUser(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["id"]

		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
	return fn
}

func createUser(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var u User
		err := decoder.Decode(&u)
		if err != nil {
			panic(err)
		}

		row, err := db.Exec(`INSERT INTO users (name, email, password, created_at) VALUES (?,?,?,?)`, u.Name, u.Email, u.Password, u.CreatedAt)
		if err != nil {
			panic(err)
		}
		var id, _ = row.LastInsertId()

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(Response{
			Status: 201,
			User: User{
				Id:        int(id),
				Name:      u.Name,
				Email:     u.Email,
				Password:  u.Password,
				CreatedAt: u.CreatedAt,
			},
		})
		if err != nil {
			return
		}
	}
	return fn
}

func updateUser(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		decoder := json.NewDecoder(r.Body)
		var data User
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)

		}
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			panic(err)
		}
		data.Id = id
		result, err := db.Exec(`UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?`, data.Name, data.Email, data.Password, data.Id)
		if err != nil {
			panic(err)
		}

		_, err = result.LastInsertId()
		if err != nil {
			panic(err)
		}

		err = json.NewEncoder(w).Encode(Response{
			Status: 200,
			User:   data,
		})
		if err != nil {
			return
		}
	}
	return fn
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	address := net.JoinHostPort(host, port)

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUrl := net.JoinHostPort(dbHost, dbPort)

	db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@("+dbUrl+")/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	usersRouter := router.PathPrefix("/api/v1").Subrouter()
	usersRouter.HandleFunc("/users", getUsers(db)).Methods("GET")
	usersRouter.HandleFunc("/users/{id}", showUser(db)).Methods("GET")
	usersRouter.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")
	usersRouter.HandleFunc("/users", createUser(db)).Methods("POST")
	usersRouter.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")

	fmt.Printf("server is running at %v \n", address)
	err = http.ListenAndServe(address, router)
	if err != nil {
		return
	}

}
