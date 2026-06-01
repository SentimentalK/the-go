/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda

References:
[1] The Go Authors. (n.d.). Package bufio. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/bufio [Accessed on: May 2026].

[2] The Go Authors. (n.d.). Package fmt. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/fmt [Accessed on: May 2026].

[3] The Go Authors. (n.d.). Package os. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/os [Accessed on: May 2026].

[4] The Go Authors. (n.d.). Package strings. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/strings [Accessed on: May 2026].
*/

package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ConsoleInput wraps a buffered reader for reusable console input.
type ConsoleInput struct {
	reader *bufio.Reader
}

// NewConsoleInput creates a console input reader that reads from standard input.
func NewConsoleInput() *ConsoleInput {
	return &ConsoleInput{
		reader: bufio.NewReader(os.Stdin),
	}
}

// GetNumber repeatedly prompts the user until a valid integer is entered.
func (input *ConsoleInput) GetNumber(prompt string) (int, error) {
	for {
		fmt.Print(prompt)

		text, err := input.reader.ReadString('\n')
		if err != nil {
			return 0, err
		}

		var number int
		if _, err := fmt.Sscanf(strings.TrimSpace(text), "%d", &number); err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		return number, nil
	}
}

// GetText prompts the user once and returns the trimmed text.
func (input *ConsoleInput) GetText(prompt string) (string, error) {
	fmt.Print(prompt)

	text, err := input.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}
