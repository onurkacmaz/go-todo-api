package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rest-api/repository"
	"rest-api/util/response"
	"strconv"
)

func GetTasks(w http.ResponseWriter, _ *http.Request) {
	response.Response{
		Status: 200,
		Data:   repository.Task{}.All(),
	}.ResponseJson(w)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task repository.Task
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		panic(err)
	}
	res := task.Create()
	response.Response{
		Status: 201,
		Data:   res,
	}.ResponseJson(w)
}

func ShowTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	task := repository.Task{Id: id}.Get()
	if task.UserId <= 0 {
		response.Response{
			Status:  404,
			Message: "Task not found.",
		}.ResponseJson(w)
		return
	}
	response.Response{
		Status: 200,
		Data:   task,
	}.ResponseJson(w)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	repository.Task{Id: id}.Delete()
	response.Response{Status: 204}.ResponseNoContent(w)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var task repository.Task
	err := decoder.Decode(&task)
	if err != nil {
		return
	}

	task.Id, _ = strconv.Atoi(params["id"])

	res := task.Update()
	if res == false {
		response.Response{
			Status:  400,
			Message: "error",
		}.ResponseJson(w)
		return
	}
	response.Response{
		Status: 200,
		Data:   task,
	}.ResponseJson(w)
	return
}
