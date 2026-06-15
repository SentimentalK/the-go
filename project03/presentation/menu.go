/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 21, 2026

References:
[1] The Go Authors. (n.d.). Package fmt. pkg.go.dev.
    [Online]. Available at https://pkg.go.dev/fmt [Accessed on: June 1, 2026].
*/

package presentation

import (
	"fmt"
	"project03/business"
	"project03/presentation/user"
)

// MenuLoop displays the console menu and routes each option to the business layer.
func MenuLoop() {
	userInput := user.NewConsoleInput()
	if _, err := business.LoadPreyRecords("./data/Prey collection & analysis - raw data.csv"); err != nil {
		fmt.Println("Error loading records:", err)
	}

	for {
		fmt.Println("Program by Xinghan Xu")
		fmt.Println("1. Reload data from dataset")
		fmt.Println("2. Save current data to new CSV file")
		fmt.Println("3. Display records")
		fmt.Println("4. Create new record")
		fmt.Println("5. Edit a record")
		fmt.Println("6. Delete a record")
		fmt.Println("7. Sort records by Lat")
		fmt.Println("8. Exit")

		choice, err := userInput.GetNumber("Please enter your choice: ")
		if err != nil {
			fmt.Println("\nExiting program.")
			return
		}

		switch choice {
		case 1:
			if _, err := business.LoadPreyRecords("./data/Prey collection & analysis - raw data.csv"); err != nil {
				fmt.Println("Error loading records:", err)
			}
		case 2:
			if err := business.SavePreyRecords(); err != nil {
				fmt.Println("Error saving records:", err)
			}
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
			business.SortPreyRecordsByLat()
			fmt.Print("Records sorted by Lat.\n\n")
		case 8:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 7.")
		}
	}
}
