package main

import (
	"bufio"
	"fmt"
	"os"
	"rest-api/database"
	"rest-api/util"
	"strings"
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
