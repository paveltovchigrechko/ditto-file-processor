package helpers

import (
	"os"
)

func CreateDir(path string) error {
	exists, err := dirExists(path)
	if err != nil {
		return err
	}

	if !exists {
		err := os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}
	return nil
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
