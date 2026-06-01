/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 21, 2026

References:
[1] The Go Authors. (n.d.). Package fmt. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/fmt [Accessed on: May 2026].
*/

package presentation

import (
	"fmt"
	"project02/business"
	"project02/presentation/user"
)

// MenuLoop displays the console menu and routes each option to the business layer.
func MenuLoop() {
	userInput := user.NewConsoleInput()

	fmt.Println("\nProgram by Xinghan Xu")
	for {
		fmt.Println("1. Reload data from dataset")
		fmt.Println("2. Save current data to new CSV file")
		fmt.Println("3. Display records")
		fmt.Println("4. Create new record")
		fmt.Println("5. Edit a record")
		fmt.Println("6. Delete a record")
		fmt.Println("7. Exit")

		choice, err := userInput.GetNumber("Please enter your choice: ")
		if err != nil {
			fmt.Println("\nExiting program.")
			return
		}

		switch choice {
		case 1:
			business.LoadPreyRecords("./data/Prey collection & analysis - raw data.csv")
		case 2:
			business.SavePreyRecords()
		case 3:
			indexes, _ := userInput.GetNumbers("Please enter record indexes to display, separated by commas: ")
			for _, index := range indexes {
				business.DisplayPreyRecords(index)
			}
		case 4:
			business.AppendPreyRecord()
		case 5:
			index, _ := userInput.GetNumber("Please enter the index of the record to edit: ")
			business.EditPreyRecord(index)
		case 6:
			index, _ := userInput.GetNumber("Please enter the index of the record to delete: ")
			business.DeletePreyRecord(index)
		case 7:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 7.")
		}
	}
}
