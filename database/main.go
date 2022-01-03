package database

import (
	"database/sql"
	"io/fs"
	"io/ioutil"
	"net"
	"path/filepath"
	"rest-api/model/config"
	"rest-api/util/file"

	_ "github.com/go-sql-driver/mysql"
)

func Instance() *sql.DB {

	c := config.Get()

	dbUrl := net.JoinHostPort(c.DbHost, c.DbPort)

	db, err := sql.Open("mysql", c.DbUser+":"+c.DbPassword+"@("+dbUrl+")/"+c.DbName+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db

}

func Migrate() {
	path := "/database/migrations/"
	fileA := file.Files(path)
	files := fileA.GetFiles()
	for _, file := range files {
		_, err := Instance().Exec(ReadFile(file, fileA.GetBasePath()))
		if err != nil {
			panic(err)
		}
	}
}

func ReadFile(file fs.FileInfo, basePath string) string {
	path := "/database/migrations/"
	q, err := ioutil.ReadFile(filepath.Join(filepath.Join(basePath, path), file.Name()))
	if err != nil {
		panic(err)
	}
	return string(q)
}
