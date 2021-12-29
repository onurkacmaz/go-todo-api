package main

import (
	"bufio"
	"fmt"
	"os"
	"rest-api/database"
	"rest-api/repository"
	"rest-api/util"
	"strings"
	"time"
)

func main() {

	cmd := os.Args[1]

	switch cmd {
	case "migrate":
		file := util.Files("/database/migrations/")
		countOfFiles := len(file.GetFiles())
		r := prompt(fmt.Sprintf("%v files found. Are you sure do you want to migrate?", countOfFiles))
		if r == "no" {
			return
		}
		database.Migrate()
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
	default:
		fmt.Print("No command provided.")
	}
}

func prompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
