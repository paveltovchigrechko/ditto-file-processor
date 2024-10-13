package ditto

import (
	"bytes"
	"encoding/json"
	"log"
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

func SplitProjectAndLocale(filename string) (string, string) {
	splitted := strings.Split(filename, nameSep)
	if len(splitted) < 3 {
		log.Printf("%s has incorrect name format, expeted: 'components__project__locale.json'\n", filename)
		return "", ""
	}

	projectName := splitted[1]
	variant := splitted[2]
	locale := defineLocale(variant)
	return projectName, locale
}

func ExtractDittoKeys(path, project string) interface{} {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Could not read %s: %s\n", path, err)
	}

	var m map[string]interface{}
	err = json.Unmarshal(f, &m)
	if err != nil {
		log.Fatalf("Could not unmarchall %s: %s\n", path, err)
	}

	if jsonBlob, ok := m[project]; !ok {
		log.Fatalf("The key '%s' was not found in %s\n", project, path)
		return ""
	} else {
		return jsonBlob
	}
}

func EncodeDittoKeys(df interface{}) []byte {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent(prefix, indent)
	err := encoder.Encode(df)
	if err != nil {
		log.Printf("Could not encode Ditto keys: %s\n", err)
		return nil
	}

	return buffer.Bytes()
}

func defineLocale(s string) string {
	switch s {
	case baseLocale:
		return defaultLocale
	default:
		return s
	}
}
