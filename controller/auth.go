package controller

func Check(email string, password string) bool {
	return IsUserExistsByCredentials(email, password)
}
