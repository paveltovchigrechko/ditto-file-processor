package ditto

import (
	"os"
	"testing"
)

func TestSplitProjectNameAndLocale(t *testing.T) {
	tests := []struct {
		input   string
		project string
		locale  string
	}{
		// All correct
		{
			"something__project__locale",
			"project",
			"locale",
		},
		// Base locale
		{
			"something__project__base.json",
			"project",
			"en.default.json",
		},
		// Base locale without extenstion
		{
			"something__project__base",
			"project",
			"base",
		},
		// Base locale with wrong extension
		{
			"something__project__base.txt",
			"project",
			"base.txt",
		},
		// One underscore
		{
			"something_project_locale",
			"",
			"",
		},
		// Three underscores
		{
			"something___project___locale",
			"_project",
			"_locale",
		},
		// Wrong delimiter
		{
			"something--project--locale",
			"",
			"",
		},
		// One part
		{
			"filename.whatever",
			"",
			"",
		},
		// Two parts
		{
			"filename__whatever",
			"",
			"",
		},
		// 4 parts
		{
			"something__project__locale__ending",
			"",
			"",
		},
		// Empty string
		{
			"",
			"",
			"",
		},
	}

	for _, test := range tests {
		p, l := SplitProjectAndLocale(test.input)

		if p != test.project {
			t.Errorf("Project name error: expected=%q, got=%q", test.project, p)
		}

		if l != test.locale {
			t.Errorf("Locale name error: expected=%q, got=%q", test.locale, l)
		}
	}
}

func TestReadDittoFiles(t *testing.T) {
	dir := "../../test/correct"
	files := ReadDittoFiles(dir)
	if files == nil {
		t.Errorf("No files were read in %s. Expected to read one file\n", dir)
	}

	if len(files) != 1 {
		t.Errorf("Wrong number of files read, expected=1, got=%d\n", len(files))
	}

	dir = "../../test/empty"
	files = ReadDittoFiles(dir)
	if files == nil {
		t.Errorf("No files were read in %s. Expected to read one file\n", dir)
	}

	if len(files) != 0 {
		t.Errorf("Wrong number of files read, expected=1, got=%d\n", len(files))
	}

	dir = "wrong_path"
	files = ReadDittoFiles(dir)
	if files != nil {
		t.Errorf("Files != nil. Expected files == nil\n")
	}
}

func TestExtractDittoKeys(t *testing.T) {
	path := "../../test/correct/test.json"
	project := "project"
	f := ExtractDittoKeys(path, project)
	if f == nil {
		t.Errorf("Error when opening file on %s. Expected no error\n", path)
	}

	wrongPath := "../../test/incorrect/test.json"
	f = ExtractDittoKeys(wrongPath, project)
	if f != nil {
		t.Errorf("Extracted file on %s is not nil, expected to be nil\n", wrongPath)
	}

	incorrectFile := "../../test/incorrect/broken.json"
	f = ExtractDittoKeys(incorrectFile, project)
	if f != nil {
		t.Errorf("Extracted incorrect file on %s is not nil, expected to be nil\n", incorrectFile)
	}

	wrongProject := "another_project"
	f = ExtractDittoKeys(path, wrongProject)
	if f != nil {
		t.Errorf("Extracted file on %s with wrong project %s is not nil, expected to be nil\n", path, wrongPath)
	}
}

func TestEncodeDittoKeys(t *testing.T) {
	text := "some text"
	encoded := EncodeDittoKeys(text)
	if encoded == nil {
		t.Errorf("Encoded text %q resulted in nil, expected not equal to nil\n", text)
	}
}

func TestCreateAndWriteJson(t *testing.T) {
	path := "../../test/correct/result.json"
	bytes := []byte{
		21, 23,
	}

	CreateAndWriteJson(path, bytes)
	defer os.Remove(path)

	_, err := os.Stat(path)
	if err != nil {
		t.Errorf("Expected to be a file on %s, got nothing\n", path)
	}
	if os.IsNotExist(err) {
		t.Errorf("Expected to be a file on %s, got nothing\n", path)
	}
}
