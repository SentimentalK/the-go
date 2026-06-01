package presentation

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ConsoleInput struct {
	reader *bufio.Reader
}

func NewConsoleInput() *ConsoleInput {
	return &ConsoleInput{
		reader: bufio.NewReader(os.Stdin),
	}
}

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

func (input *ConsoleInput) GetText(prompt string) (string, error) {
	fmt.Print(prompt)

	text, err := input.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}
