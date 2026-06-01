/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 21, 2026

References:
[1] The Go Authors. (n.d.). Package filepath. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/path/filepath [Accessed on: May 2026].

[2] The Go Authors. (n.d.). Package regexp. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/regexp [Accessed on: May 2026].

[3] The Go Authors. (n.d.). Package testing. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/testing [Accessed on: May 2026].
*/

package tests

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// TestGoFilesHaveRequiredReferenceHeaders checks that every Go file has the
// required course header and at least one IEEE-style website reference.
func TestGoFilesHaveRequiredReferenceHeaders(t *testing.T) {
	projectRoot := filepath.Clean("..")

	// These fields must appear in each file's top comment block.
	requiredFields := []string{
		"Author: Xinghan Xu",
		"Course: CST8002 Programming Language Research Project",
		"Professor: Stanley Pieda",
		"Due Date: June 21, 2026",
		"References:",
	}

	// Matches the required website reference format:
	// Author. (date). Title. domain. [online]. Available at URL [Accessed on: date].
	referencePattern := regexp.MustCompile(`(?m)^\[\d+\] .+\. \((n\.d\.|\d{4}(, [A-Z][a-z]{2}\. \d{1,2})?)\)\. .+\. [A-Za-z0-9.-]+\.\n\s+\[online\]\. Available at https?://\S+ \[Accessed on: .+\]\.$`)

	err := filepath.WalkDir(projectRoot, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Only files are checked; directories are skipped during traversal.
		if entry.IsDir() {
			if entry.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}

		// Check all Go files, including test files.
		if filepath.Ext(path) != ".go" {
			return nil
		}

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		content := string(contentBytes)

		// Report every missing field so one test run shows all header problems.
		for _, field := range requiredFields {
			if !strings.Contains(content, field) {
				t.Errorf("%s is missing %q", path, field)
			}
		}
		if !referencePattern.MatchString(content) {
			t.Errorf("%s does not include a reference in the required IEEE website format", path)
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
