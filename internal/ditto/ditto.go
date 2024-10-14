package ditto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

const (
	nameSep      string = "__"
	extensionSep string = "."
	prefix       string = ""
	indent       string = "    "

	baseLocale    string = "base.json"
	defaultLocale string = "en.default.json"
)

func ReadDittoFiles(dir string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func SplitProjectAndLocale(filename string) (string, string, error) {
	splitted := strings.Split(filename, nameSep)
	if len(splitted) != 3 {
		return "", "", fmt.Errorf("%s has incorrect name format, expected: 'components__project__locale.json'", filename)
	}

	projectName := splitted[1]
	variant := splitted[2]
	locale := defineLocale(variant)
	return projectName, locale, nil
}

func ExtractDittoKeys(path, project string) (interface{}, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %s", path, err)
	}

	var m map[string]interface{}
	err = json.Unmarshal(f, &m)
	if err != nil {
		return nil, fmt.Errorf("could not unmarchall %s: %s", path, err)
	}

	if jsonBlob, ok := m[project]; !ok {
		return nil, fmt.Errorf("the key '%s' was not found in %s", project, path)
	} else {
		return &jsonBlob, nil
	}
}

func EncodeDittoKeys(df interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent(prefix, indent)

	err := encoder.Encode(df)
	if err != nil {
		return nil, fmt.Errorf("could not encode Ditto keys: %s", err)
	}
	return buffer.Bytes(), nil
}

func CreateAndWriteJson(path string, encoded []byte) error {
	newFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create file %s: %s", path, err)
	}

	_, err = newFile.Write(encoded)
	if err != nil {
		return fmt.Errorf("could not write to file %s: %s", path, err)
	}
	return nil
}

func defineLocale(s string) string {
	switch s {
	case baseLocale:
		return defaultLocale
	default:
		return s
	}
}
