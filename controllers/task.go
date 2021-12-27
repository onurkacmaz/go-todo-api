package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Task struct {
	Id        int
	UserId    int
	Title     string
	Content   string
	Status    string
	CreatedAt string
}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	var tasks []Task

	rows, err := db.Query(`SELECT * FROM tasks ORDER BY created_at DESC`)
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
		var t Task
		err := rows.Scan(&t.Id, &t.UserId, &t.Title, &t.Content, &t.Status, &t.CreatedAt)
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, t)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status int
		Tasks  []Task
	}{
		200,
		tasks,
	})
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

}

func ShowTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	rows, err := db.Query(`SELECT * FROM tasks WHERE id = ?`, params["id"])
	if err != nil {
		panic(err)
	}
	count := 0
	var task Task
	for rows.Next() {
		err := rows.Scan(&task.Id, &task.UserId, &task.Title, &task.Content, &task.Status, &task.CreatedAt)
		if err != nil {
			panic(err)
		}
		count++
	}
	if count <= 0 {
		json.NewEncoder(w).Encode(struct {
			Status int
			Task   []string
		}{
			Status: 200,
		})
		return
	}
	json.NewEncoder(w).Encode(struct {
		Status int
		Task   Task
	}{
		Status: 200,
		Task:   task,
	})
	return
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

}
