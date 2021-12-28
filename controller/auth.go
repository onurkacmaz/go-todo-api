package controller

func Check(email string, password string) bool {
	return GetUserByCredentials(email, password)
}
