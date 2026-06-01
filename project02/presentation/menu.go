package presentation

import (
	"bufio"
	"fmt"
	"os"
	"project02/business"
	"strings"
)

func GetUserInput(reader *bufio.Reader) (int, error) {
	for {
		fmt.Print("Please enter a number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}

		var n int
		if _, err := fmt.Sscanf(strings.TrimSpace(input), "%d", &n); err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		return n, nil
	}
}

func MenuLoop() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nProgram by Xinghan Xu")
		fmt.Println("1. Reload data from dataset")
		fmt.Println("2. Save current data to new CSV file")
		fmt.Println("3. Display records")
		fmt.Println("4. Create new record")
		fmt.Println("5. Edit a record")
		fmt.Println("6. Delete a record")
		fmt.Println("7. Exit")

		choice, err := GetUserInput(reader)
		if err != nil {
			fmt.Println("\nExiting program.")
			return
		}

		var records []business.PreyRecord

		switch choice {
		case 1:
			records, _ = business.LoadPreyRecords("./data/Prey collection & analysis - raw data.csv")
		case 2:
			business.SavePreyRecords(records)
		case 3:
			index, _ := GetUserInput(reader)
			business.DisplayPreyRecords(index, records)
		case 4:
			business.CreatePreyRecord()
		case 5:
			index, _ := GetUserInput(reader)
			business.EditPreyRecord(index)
		case 6:
			index, _ := GetUserInput(reader)
			business.DeletePreyRecord(&records, index)
		case 7:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 7.")
		}
	}
}
