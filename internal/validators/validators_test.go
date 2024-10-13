package validators

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

const (
	project = "project"
	locale  = "locale"
	file    = "file"
)

func TestValidateFiles(t *testing.T) {
	dirWithFiles := "../../input"
	files, _ := os.ReadDir(dirWithFiles)
	err := ValidateFiles(files, dirWithFiles)
	if err != nil {
		t.Errorf("Error when validating files in %q: expected=nil, got=%s", dirWithFiles, err)
	}

	curDir := "."
	files, _ = os.ReadDir(curDir)
	err = ValidateFiles(files, curDir)
	if err != nil {
		t.Errorf("Error when validating files in %q: expected=nil, got=%s", curDir, err)
	}

	emptyDir := "../../test/empty"
	files, _ = os.ReadDir(emptyDir)
	err = ValidateFiles(files, emptyDir)
	expected := fmt.Errorf("no files in %s", emptyDir)
	if errors.Is(err, expected) {
		t.Errorf("Error when validating files in %q: expected=%q, got=%q", emptyDir, expected, err)
	}
}

func TestValidateNames(t *testing.T) {
	tests := []struct {
		project  string
		locale   string
		fn       string
		expected error
	}{
		// All correct
		{
			project,
			locale,
			file,
			nil,
		},
		// Project empty
		{
			"",
			locale,
			file,
			fmt.Errorf("project name was not parsed for %s", file),
		},
		// Locale empty
		{
			project,
			"",
			file,
			fmt.Errorf("locale was not parsed for %s", file),
		},
		// File empty
		{
			project,
			locale,
			"",
			nil,
		},
		// Project and locale empty
		{
			"",
			"",
			file,
			fmt.Errorf("project name was not parsed for %s", file),
		},
		// Project and file empty
		{
			"",
			locale,
			"",
			fmt.Errorf("project name was not parsed for %s", ""),
		},
		// Locale and file empty
		{
			project,
			"",
			"",
			fmt.Errorf("locale name was not parsed for %s", ""),
		},
		// All empty
		{
			"",
			"",
			"",
			fmt.Errorf("project name was not parsed for %s", ""),
		},
	}

	for _, test := range tests {
		result := ValidateNames(test.project, test.locale, test.fn)
		if !errors.Is(test.expected, result) {
			t.Errorf("Error when validating names: expected=%q, got=%q", test.expected, result)
		}
	}
}
