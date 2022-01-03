package controller

import (
	"encoding/json"
	"net/http"
	"rest-api/repository"
	"rest-api/util/response"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	response.Response{Status: 200, Data: repository.User{}.All()}.ResponseJson(w)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	user := repository.User{Id: id}.Get()
	if user.Email == "" {
		response.Response{Status: 404, Data: nil}.ResponseJson(w)
		return
	}
	response.Response{Status: 200, Data: user}.ResponseJson(w)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var u repository.User
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	u.Id = id
	u.Delete()
	response.Response{Status: 204, Data: nil}.ResponseNoContent(w)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u repository.User
	err := decoder.Decode(&u)
	if err != nil {
		panic(err.Error())
	}
	res := u.Create()
	response.Response{Status: 201, Data: res}.ResponseJson(w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)

	var user repository.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err.Error())
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	user.Id = id

	res := user.Update()
	if res == false {
		response.Response{Status: 400, Data: nil, Message: "error"}.ResponseJson(w)
		return
	}
	response.Response{Status: 200, Data: res}.ResponseJson(w)
}

func IsUserExistsByCredentials(email string, password string) bool {
	user := repository.User{Email: email, Password: password}.GetByCredentials()
	if user.Id > 0 {
		return true
	}
	return false
}
