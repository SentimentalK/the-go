/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 21, 2026

References:
[1] The Go Authors. (n.d.). Package bufio. pkg.go.dev.
    [Online]. Available at https://pkg.go.dev/bufio [Accessed on: June 1, 2026].

[2] The Go Authors. (n.d.). Package fmt. pkg.go.dev.
    [Online]. Available at https://pkg.go.dev/fmt [Accessed on: June 1, 2026].

[3] The Go Authors. (n.d.). Package os. pkg.go.dev.
    [Online]. Available at https://pkg.go.dev/os [Accessed on: June 1, 2026].

[4] The Go Authors. (n.d.). Package strings. pkg.go.dev.
    [Online]. Available at https://pkg.go.dev/strings [Accessed on: June 1, 2026].

[5] The Go Authors. (n.d.). Package strconv. pkg.go.dev.
    [Online]. Available at https://pkg.go.dev/strconv [Accessed on: June 1, 2026].
*/

package user

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

// GetNumbers prompts the user for one or more comma-separated integers.
func (input *ConsoleInput) GetNumbers(prompt string) ([]int, error) {
	for {
		text, err := input.GetText(prompt)
		if err != nil {
			return nil, err
		}

		parts := strings.Split(text, ",")
		numbers := []int{}
		valid := true

		for _, part := range parts {
			number, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				fmt.Println("Invalid input. Please enter numbers separated by commas.")
				valid = false
				break
			}
			numbers = append(numbers, number)
		}

		if valid {
			return numbers, nil
		}
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
