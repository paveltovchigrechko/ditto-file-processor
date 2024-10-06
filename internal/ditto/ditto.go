package ditto

import (
	"strings"
)

const (
	NAME_SEP      string = "__"
	EXTENSION_SEP string = "."
	BASE_LOCALE   string = "base.json"
	DEF_LOCALE    string = "en.default.json"
)

func SplitProjectAndLocale(filename string) (string, string) {
	splitted := strings.Split(filename, NAME_SEP)
	// TODO handle the case with wrong file: slice with 1 element
	// if len(splitted) < 3 {
	// Do something
	// }
	projectName := splitted[1]
	variant := splitted[2]
	locale := defineLocale(variant)
	return projectName, locale
}

func defineLocale(s string) string {
	switch s {
	case BASE_LOCALE:
		return DEF_LOCALE
	default:
		return s
	}
}
