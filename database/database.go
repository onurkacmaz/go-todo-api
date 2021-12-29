package database

import (
	"database/sql"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"rest-api/config"
	"rest-api/util"

	_ "github.com/go-sql-driver/mysql"
)

func Instance() *sql.DB {

	config.LoadLocalEnv()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUrl := net.JoinHostPort(dbHost, dbPort)

	db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@("+dbUrl+")/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db

}

func Migrate() {
	path := "/database/migrations/"
	file := util.Files(path)
	files := file.GetFiles()
	basePath := file.GetBasePath()
	for _, file := range files {
		q, err := ioutil.ReadFile(filepath.Join(filepath.Join(basePath, path), file.Name()))
		if err != nil {
			panic(err)
		}

		defer func(q string) {
			_, err = Instance().Exec(string(q))
			if err != nil {
				panic(err)
			}
		}(string(q))
	}
}
