package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"rest-api/database"
	"strconv"

	"github.com/gorilla/mux"
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

var db = database.Instance()

func GetUsers(w http.ResponseWriter, r *http.Request) {

	type Response struct {
		Status int
		Users  []User
	}

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Status: 200,
		Users:  users,
	})

}

func ShowUser(w http.ResponseWriter, r *http.Request) {
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
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
