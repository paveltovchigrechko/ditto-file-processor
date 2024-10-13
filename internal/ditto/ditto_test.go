package ditto

import "testing"

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
