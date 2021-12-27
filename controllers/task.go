package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Task struct {
	Id        int
	UserId    int
	Title     string
	Content   string
	Status    string
	CreatedAt string
}

func GetTasks(w http.ResponseWriter, _ *http.Request) {

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
	err = json.NewEncoder(w).Encode(struct {
		Status int
		Tasks  []Task
	}{
		200,
		tasks,
	})
	if err != nil {
		panic(err)
	}
	return
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task Task

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		panic(err)
	}

	row, err := db.Exec(
		`INSERT INTO tasks (user_id, title, content, status, created_at) VALUES (?,?,?,?,?)`,
		task.UserId, task.Title, task.Content, task.Status, task.CreatedAt,
	)
	if err != nil {
		panic(err)
	}

	id, err := row.LastInsertId()
	if err != nil {
		panic(err)
	}

	task.Id = int(id)

	err = json.NewEncoder(w).Encode(struct {
		Status int
		Task   Task
	}{
		Status: 201,
		Task:   task,
	})
	if err != nil {
		panic(err)
	}
	return

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
		err = json.NewEncoder(w).Encode(struct {
			Status int
			Task   []string
		}{
			Status: 200,
		})
		if err != nil {
			panic(err)
		}
		return
	}
	err = json.NewEncoder(w).Encode(struct {
		Status int
		Task   Task
	}{
		Status: 200,
		Task:   task,
	})
	if err != nil {
		panic(err)
	}
	return
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	_, err := db.Exec(`DELETE FROM tasks WHERE id = ?`, id)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var task Task
	err := decoder.Decode(&task)
	if err != nil {
		return
	}

	task.Id, _ = strconv.Atoi(params["id"])

	_, err = db.Exec(
		`UPDATE tasks SET title = ?, content = ?, status = ?, created_at = ? WHERE id = ? AND user_id = ?`,
		task.Title, task.Content, task.Status, task.CreatedAt, task.Id, task.UserId,
	)

	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(w).Encode(struct {
		Status int
		Task   Task
	}{
		Status: 200,
		Task:   task,
	})
	if err != nil {
		panic(err)
	}
	return
}
