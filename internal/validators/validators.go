package validators

import (
	"fmt"
	"io/fs"
)

func ValidateFiles(files []fs.DirEntry, dir string) error {
	if len(files) == 0 {
		return fmt.Errorf("no files in %s", dir)
	}
	return nil
}

func ValidateNames(project, locale, fn string) error {
	if project == "" {
		return fmt.Errorf("project name was not parsed for %s", fn)
	} else if locale == "" {
		return fmt.Errorf("locale was not parsed for %s", fn)
	}

	return nil
}
