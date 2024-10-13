package helpers

import (
	"fmt"
	"log"
	"os"
)

func CreateDir(path string) {
	dirExists, _ := dirExists(path)

	if !dirExists {
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Println(err)
		}
	}
}

func ReadArgs() []string {
	args := os.Args[1:]
	fmt.Println(args)
	if len(args) != 1 {
		log.Fatalf("Wrong number of arguments: expected only one file name")
	}

	return args
}

func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
