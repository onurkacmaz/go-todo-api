package repository

import (
	"database/sql"
	"rest-api/database"
)

type Task struct {
	Id        int
	UserId    int
	Title     string
	Content   string
	Status    string
	CreatedAt string
}

func (t Task) Create() Task {
	row, err := database.Instance().Exec(`INSERT INTO tasks (user_id, title, content, status, created_at) VALUES (?,?,?,?,?)`, t.UserId, t.Title, t.Content, t.Status, t.CreatedAt)
	if err != nil {
		panic(err)
	}
	var id, _ = row.LastInsertId()
	t.Id = int(id)
	return t
}

func (t Task) Update() bool {
	result, err := database.Instance().Exec(`UPDATE tasks SET user_id = ?, title = ?, content = ?, status = ?, created_at = ? WHERE id = ?`, t.UserId, t.Title, t.Content, t.Status, t.CreatedAt, t.Id)
	if err != nil {
		panic(err)
	}

	res, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	if res == 0 {
		return false
	}
	return true
}

func (t Task) Delete() {
	_, err := database.Instance().Exec(`DELETE FROM tasks WHERE id = ?`, t.Id)
	if err != nil {
		panic(err)
	}
	return
}

func (t Task) Get() Task {
	var task Task
	rows, err := database.Instance().Query(`SELECT * FROM tasks WHERE id = ?`, t.Id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&task.Id, &task.UserId, &task.Title, &task.Content, &task.Status, &task.CreatedAt)
		if err != nil {
			panic(err)
		}
	}
	return task
}

func (t Task) All() []Task {

	var tasks []Task

	rows, err := database.Instance().Query(`SELECT * FROM tasks ORDER BY created_at DESC`)
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
	return tasks
}
