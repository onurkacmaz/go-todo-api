package repository

import (
	"database/sql"
	"rest-api/database"
)

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt string
}

func (u User) Create() User {
	row, err := database.Instance().Exec(`INSERT INTO users (name, email, password, created_at) VALUES (?,?,?,?)`, u.Name, u.Email, u.Password, u.CreatedAt)
	if err != nil {
		panic(err)
	}
	var id, _ = row.LastInsertId()
	u.Id = int(id)
	return u
}

func (u User) Update() bool {
	result, err := database.Instance().Exec(`UPDATE users SET name = ?, email = ?, password = ?, created_at = ? WHERE id = ?`, u.Name, u.Email, u.Password, u.CreatedAt, u.Id)
	if err != nil {
		panic(err)
	}

	res, err := result.RowsAffected()
	if err != nil || res == 0 {
		return false
	}
	return true
}

func (u User) Delete() {
	_, err := database.Instance().Exec(`DELETE FROM users WHERE id = ?`, u.Id)
	if err != nil {
		panic(err)
	}
	return
}

func (u User) Get() User {
	var user User
	rows, err := database.Instance().Query(`SELECT * FROM users WHERE id = ?`, u.Id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			panic(err)
		}
	}
	return user
}

func (u User) GetByCredentials() User {
	var user User
	rows, err := database.Instance().Query(`SELECT * FROM users WHERE email = ? AND password = ?`, u.Email, u.Password)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			panic(err)
		}
	}
	return user
}

func (u User) All() []User {

	var users []User

	rows, err := database.Instance().Query(`SELECT * FROM users ORDER BY created_at DESC`)
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
	return users
}
