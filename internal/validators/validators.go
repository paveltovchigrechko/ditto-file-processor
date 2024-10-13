package validators

import (
	"errors"
	"fmt"
	"os"
)

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
