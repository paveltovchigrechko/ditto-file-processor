package ditto

import (
	"io/fs"
	"os"
)

type OSProvider interface {
	ReadDir(string) ([]fs.DirEntry, error)
}

type OSWrapper struct{}

func (osw OSWrapper) ReadDir(dir string) ([]fs.DirEntry, error) {
	return os.ReadDir(dir)
}
