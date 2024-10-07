package validators

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func ValidateFiles(files []fs.DirEntry, dir string) error {
	if files == nil {
		return fmt.Errorf("no files in %s", dir)
	}
	return nil
}

func ValidateFile(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}

func ValidateNames(project, locale, fn string) error {
	if project == "" {
		return fmt.Errorf("project name was not parsed for '%s'", fn)
	} else if locale == "" {
		return fmt.Errorf("locale was not parsed for '%s'", fn)
	}

	return nil
}

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
