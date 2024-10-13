package helpers

import (
	"log"
	"os"
)

func CreateDir(path string) {
	exists, err := dirExists(path)
	if err != nil {
		log.Printf("Error when defining if the directory %s exists: %s\n", path, err)
	}

	if !exists {
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Println(err)
		}
	}
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
