package main

import (
	"bufio"
	"fmt"
	"os"
	main2 "rest-api/database"
	"rest-api/repository"
	"rest-api/util/file"
	"strconv"
	"strings"
	"time"
)

func main() {

	cmd := os.Args[1]

	switch cmd {
	case "migrate":
		file := file.Files("/database/migrations/")
		countOfFiles := len(file.GetFiles())
		r := prompt(fmt.Sprintf("%v files found. Are you sure do you want to migrate?", countOfFiles))
		if r == "no" {
			return
		}
		main2.Migrate()
		fmt.Println("Tables migrated successfully.")
	case "user":
		if os.Args[2] == "create" {
			name := prompt("Name:")
			email := prompt("Email:")
			password := prompt("Password:")
			res := repository.User{
				Name:      name,
				Email:     email,
				Password:  password,
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			}.Create()
			if res.Id > 0 {
				fmt.Println("User created successfully")
			}
		}
		break
	case "token":
		if os.Args[2] == "create" {
			userId, _ := strconv.Atoi(prompt("User Id:"))
			expiredAt := prompt("Expired At:")
			res := repository.Token{
				UserId:    userId,
				Token:     repository.GenerateToken(10),
				ExpiredAt: expiredAt,
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			}.Create()
			if res.Id > 0 {
				fmt.Println("Token generated and created successfully")
			}
		}
		break
	default:
		fmt.Print("No command provided.")
	}
}

func prompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		_, err := fmt.Fprint(os.Stderr, label+" ")
		if err != nil {
			return ""
		}
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
