/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project

References:
[1] The Go Authors. (n.d.). Package os. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/os [Accessed on: May 25, 2026].

[2] The Go Authors. (n.d.). Package filepath. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/path/filepath [Accessed on: May 25, 2026].
*/

package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

const sourceFolder = "src"

// ResolveFilePath returns a path inside the src folder for the given file name.
func ResolveFilePath(fileName string) (string, error) {
	sourceFolderInfo, err := os.Stat(sourceFolder)
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("src folder does not exist")
		}

		return "", err
	}
	if !sourceFolderInfo.IsDir() {
		return "", errors.New("src is not a folder")
	}

	cleanName := strings.TrimPrefix(filepath.Clean("/"+strings.TrimSpace(fileName)), string(filepath.Separator))
	if cleanName == "." || cleanName == "" {
		return "", errors.New("file name is required")
	}
	if filepath.Ext(cleanName) != ".txt" {
		return "", errors.New("file must have a .txt extension")
	}

	return filepath.Join(sourceFolder, cleanName), nil
}

// ReadData returns the current content of a file inside the src folder.
func ReadData(fileName string) (string, error) {
	filePath, err := ResolveFilePath(fileName)
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("file does not exist")
		}

		return "", err
	}

	return string(content), nil
}

// AppendData adds a line of text to a file inside the src folder.
func WriteData(fileName string, text string) error {
	filePath, err := ResolveFilePath(fileName)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text + "\n")
	return err
}
