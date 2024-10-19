package ditto

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

type mockReader struct{}

func (mr mockReader) ReadDir(string) ([]fs.DirEntry, error) {
	return []fs.DirEntry{}, nil
}

func TestSplitProjectNameAndLocale(t *testing.T) {
	tests := []struct {
		input   string
		project string
		locale  string
		err     error
	}{
		// All correct
		{
			input:   "something__project__locale",
			project: "project",
			locale:  "locale",
			err:     nil,
		},
		// Base locale
		{
			input:   "something__project__base.json",
			project: "project",
			locale:  "en.default.json",
			err:     nil,
		},
		// Base locale without extenstion
		{
			input:   "something__project__base",
			project: "project",
			locale:  "base",
			err:     nil,
		},
		// Base locale with wrong extension
		{
			input:   "something__project__base.txt",
			project: "project",
			locale:  "base.txt",
			err:     nil,
		},
		// One underscore
		{
			input:   "something_project_locale",
			project: "",
			locale:  "",
			err:     fmt.Errorf("something_project_locale has incorrect name format, expected: 'components__project__locale.json'"),
		},
		// Three underscores
		{
			input:   "something___project___locale",
			project: "_project",
			locale:  "_locale",
			err:     nil,
		},
		// Wrong delimiter
		{
			input:   "something--project--locale",
			project: "",
			locale:  "",
			err:     fmt.Errorf("something--project--locale has incorrect name format, expected: 'components__project__locale.json'"),
		},
		// One part
		{
			input:   "filename.whatever",
			project: "",
			locale:  "",
			err:     fmt.Errorf("filename.whatever has incorrect name format, expected: 'components__project__locale.json'"),
		},
		// Two parts
		{
			input:   "filename__whatever",
			project: "",
			locale:  "",
			err:     fmt.Errorf("filename__whatever has incorrect name format, expected: 'components__project__locale.json'"),
		},
		// 4 parts
		{
			input:   "something__project__locale__ending",
			project: "",
			locale:  "",
			err:     fmt.Errorf("something__project__locale__ending has incorrect name format, expected: 'components__project__locale.json'"),
		},
		// Empty string
		{
			input:   "",
			project: "",
			locale:  "",
			err:     fmt.Errorf(" has incorrect name format, expected: 'components__project__locale.json'"),
		},
	}

	for _, test := range tests {
		p, l, err := SplitProjectAndLocale(test.input)

		if p != test.project {
			t.Errorf("Project name error: expected=%q, got=%q", test.project, p)
		}

		if l != test.locale {
			t.Errorf("Locale name error: expected=%q, got=%q", test.locale, l)
		}

		if !errors.Is(err, test.err) {
			t.Errorf("Error %q is not equal to expected %q", err, test.err)
		}
	}
}

func TestReadDittoFilesCorrect(t *testing.T) {
	// Split scenarios into 3 tests (correct / empty / incorrect)
	// Alternative: mockup file system interface
	// 1. Create structure DittoHelper
	// 2. Make all ditto functions -> new struct methods
	// 3. Define interface{} OSProvider with ReadDir method
	// 4. Add a DittoHelper structure field with OSProvider type
	// 5. Inside ReadDittoFile call OSProvider instead of std.os.ReadDir
	dh := New(OSWrapper{})
	dir := "../../test/correct"
	files, err := dh.ReadDittoFiles(dir)

	// expected := mockReader{}
	if files == nil {
		t.Errorf("No files were read in %s. Expected to read one file\n", dir)
	}
	if len(files) != 1 {
		t.Errorf("Wrong number of files read, expected=1, got=%d\n", len(files))
	}
	if err != nil {
		t.Errorf("Error while reading directory: %s. Expected error equal nil\n", err)
	}
}

func TestReadDittoFilesEmpty(t *testing.T) {
	dh := New(OSWrapper{})
	dir := "../../test/empty"
	files, err := dh.ReadDittoFiles(dir)
	if files == nil {
		t.Errorf("No files were read in %s. Expected to read one file\n", dir)
	}
	if len(files) != 0 {
		t.Errorf("Wrong number of files read, expected=1, got=%d\n", len(files))
	}
	if err != nil {
		t.Errorf("Error while reading directory: %s. Expected error equal nil\n", err)
	}
}

func TestReadDittoFilesWrong(t *testing.T) {
	dh := New(OSWrapper{})
	dir := "wrong_path"
	files, err := dh.ReadDittoFiles(dir)
	if files != nil {
		t.Errorf("Files != nil. Expected files == nil\n")
	}
	if err == nil {
		t.Errorf("Error is nil, expected error not equal to nil\n")
	}
}

func TestExtractDittoKeys(t *testing.T) {
	path := "../../test/correct/test.json"
	project := "project"
	f, _ := ExtractDittoKeys(path, project)
	if f == nil {
		t.Errorf("Error when opening file on %s. Expected no error\n", path)
	}

	wrongPath := "../../test/incorrect/test.json"
	f, _ = ExtractDittoKeys(wrongPath, project)
	if f != nil {
		t.Errorf("Extracted file on %s is not nil, expected to be nil\n", wrongPath)
	}

	incorrectFile := "../../test/incorrect/broken.json"
	f, _ = ExtractDittoKeys(incorrectFile, project)
	if f != nil {
		t.Errorf("Extracted incorrect file on %s is not nil, expected to be nil\n", incorrectFile)
	}

	wrongProject := "another_project"
	f, _ = ExtractDittoKeys(path, wrongProject)
	if f != nil {
		t.Errorf("Extracted file on %s with wrong project %s is not nil, expected to be nil\n", path, wrongPath)
	}
}

func TestEncodeDittoKeys(t *testing.T) {
	// 1. Test negative scenario with err != nil, see json.Marchall
	text := "some text"
	encoded, _ := EncodeDittoKeys(text)
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
