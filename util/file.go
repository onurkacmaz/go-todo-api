package util

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileRes struct {
	Files    []fs.FileInfo
	BasePath string
}

func (f *FileRes) GetBasePath() string {
	return f.BasePath
}

func (f *FileRes) GetFiles() []fs.FileInfo {
	return f.Files
}

func Files(path string) FileRes {

	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir, err := os.Open(dirname)
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(filepath.Join(dir.Name(), path))
	if err != nil {
		panic(err)
	}

	return FileRes{
		Files:    files,
		BasePath: dir.Name(),
	}
}
