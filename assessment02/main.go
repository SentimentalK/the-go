/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project

References:
[1] The Go Authors, “Package fmt,” Go Packages.
    [Online]. Available: https://pkg.go.dev/fmt.
    [Accessed: May 25, 2026].

[2] The Go Authors, “Package bufio,” Go Packages.
    [Online]. Available: https://pkg.go.dev/bufio.
    [Accessed: May 25, 2026].

[3] The Go Authors, “Package os,” Go Packages.
    [Online]. Available: https://pkg.go.dev/os.
    [Accessed: May 25, 2026].

[4] The Go Authors, “Package strings,” Go Packages.
    [Online]. Available: https://pkg.go.dev/strings.
    [Accessed: May 25, 2026].

[5] The Go Authors, “How to Write Go Code,” The Go Programming Language.
    [Online]. Available: https://go.dev/doc/code.
    [Accessed: May 25, 2026].
*/

// Package main demonstrates a small console program written in Go.
package main

import (
	"assessment02/file"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const studentName = "Xinghan Xu"

// main is the entry point of the program.
func main() {
	menuLoop()
}

// menuLoop displays a simple menu and handles user input.
func menuLoop() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nProgram by Xinghan Xu")
		fmt.Println("1. Read file")
		fmt.Println("2. Write file")
		fmt.Println("3. Exit program")
		fmt.Print("Please choose an option: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("\nExiting program.")
			return
		}

		var choice int
		if _, err := fmt.Sscanf(strings.TrimSpace(input), "%d", &choice); err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter file name: ")
			fileName, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading file name:", err)
				break
			}

			content, err := file.ReadData(fileName)
			if err != nil {
				fmt.Println("Error reading file:", err)
				break
			}

			if strings.TrimSpace(content) == "" {
				fmt.Println("The file is empty.")
			} else {
				fmt.Println("File content:")
				fmt.Print(content)
			}
		case 2:
			fmt.Print("Enter text to write: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				break
			}

			fmt.Print("Enter file name: ")
			fileName, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading file name:", err)
				break
			}

			if err := file.WriteData(fileName, strings.TrimSpace(text)); err != nil {
				fmt.Println("Error writing file:", err)
				break
			}

			fmt.Println("Text written to file.")
		case 3:
			fmt.Println("Exiting program.")
			return
		}

	}
}
