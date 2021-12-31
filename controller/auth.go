package controller

import (
	"encoding/json"
	"net/http"
	"rest-api/repository"
	"rest-api/util"
	"time"
)

type Credentials struct {
	Email    string
	Password string
}

func Check(email string, password string) bool {
	return IsUserExistsByCredentials(email, password)
}

func IsTokenValid(token string) bool {
	t := repository.Token{Token: token}.Get()
	if t.Id <= 0 {
		return false
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	expiredAt, _ := time.Parse(time.RFC3339, t.ExpiredAt)
	formattedExpiredAt := expiredAt.Format("2006-01-02 15:04:05")

	return !(now >= formattedExpiredAt)
}

func CreateToken(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var credentials Credentials
	err := decoder.Decode(&credentials)
	if err != nil {
		panic(err)
	}

	isUserExists := Check(credentials.Email, credentials.Password)
	if !isUserExists {
		util.Response{
			Status:  401,
			Message: "Invalid Credentials",
		}.ResponseJson(w)
		return
	}

	u := repository.User{Email: credentials.Email, Password: credentials.Password}.GetByCredentials()
	if u.Id <= 0 {
		util.Response{
			Status:  401,
			Message: "Invalid Credentials",
		}.ResponseJson(w)
		return
	}

	util.Response{
		Status: 201,
		Data:   repository.Token{}.Create(u),
	}.ResponseJson(w)
}
