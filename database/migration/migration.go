package migration

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"rest-api/database"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func GetFiles() []fs.FileInfo {
	files, err := ioutil.ReadDir(basepath + "/sql/")
	if err != nil {
		panic(err)
	}
	return files
}

func Migrate() {
	for _, file := range GetFiles() {
		query, err := ioutil.ReadFile(filepath.Join(basepath+"/sql/", file.Name()))
		if err != nil {
			panic(err)
		}
		db := database.Instance()
		defer func(query string) {
			db.Exec(query)
		}(string(query))
	}

}
