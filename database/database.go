package database

import (
	"database/sql"
	"io/fs"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"rest-api/config"
	"runtime"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
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

func GetFiles() []fs.FileInfo {
	files, err := ioutil.ReadDir(basePath + "/migrations/")
	if err != nil {
		panic(err)
	}
	return files
}

func Migrate() {
	for _, file := range GetFiles() {
		byte, err := ioutil.ReadFile(filepath.Join(basePath+"/migrations/", file.Name()))
		if err != nil {
			panic(err)
		}
		query := strings.Replace(string(byte), "{{DB_NAME}}", os.Getenv("DB_NAME"), -1)
		defer func(query string) {
			Instance().Exec(query)
		}(query)
	}
}
