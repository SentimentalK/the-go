/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda

References:
[1] The Go Authors. (n.d.). Package fmt. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/fmt [Accessed on: May 2026].

[2] Fisheries and Oceans Canada. (2024, Dec. 16). Spatiotemporal variation in anadromous Arctic char (Salvelinus alpinus) foraging ecology and its influence on muscle pigmentation along western Hudson Bay, Nunavut, Canada. open.canada.ca.
    [online]. Available at https://open.canada.ca/data/en/dataset/9cbcf710-a2a1-11ef-8ccf-55cc7f028297 [Accessed on: Apr. 30, 2026].
*/

package business

import (
	"fmt"
	"project02/model"
	"project02/persistence"
	"project02/presentation/user"
)

// Prey stores the current in-memory collection of prey records.
type Prey struct {
	Records []model.PreyRecord
}

var prey = &Prey{}
var userInput = user.NewConsoleInput()

// LoadPreyRecords loads prey records from the CSV file into memory.
func LoadPreyRecords(filePath string) ([]model.PreyRecord, error) {
	records, err := persistence.ReadLinesFor(100, filePath)
	if err != nil {
		return nil, err
	}

	prey.Records = records
	fmt.Printf("%d records loaded.\n\n", len(records))
	return records, nil
}

// SavePreyRecords writes the current in-memory records to a new CSV file.
func SavePreyRecords() error {
	if prey.Records == nil {
		return fmt.Errorf("no records to save")
	}

	return persistence.WriteToCSV(prey.Records)
}

// DisplayPreyRecords prints one prey record by its one-based index.
func DisplayPreyRecords(index int) {

	if index < 1 || index > len(prey.Records) {
		fmt.Printf("Invalid index. Please enter a number between 1 and %d.\n", len(prey.Records))
		return
	}

	record := prey.Records[index-1]
	fmt.Printf("\n--Record %d:--\n", index)
	fmt.Printf("Year: %s\n", record.Year)
	fmt.Printf("Species: %s\n", record.Species)
	fmt.Printf("Common Name: %s\n", record.CommonName)
	fmt.Printf("Study Site: %s\n", record.StudySite)
	fmt.Printf("Associated Community: %s\n", record.AssociatedCommunity)
	fmt.Printf("Retinol: %s\n---\n\n", record.Retinol)
}

// CreatePreyRecord asks the user for record values and returns a new prey record.
func CreatePreyRecord() model.PreyRecord {
	var record model.PreyRecord

	record.Year, _ = userInput.GetText("Enter Year: ")
	record.Species, _ = userInput.GetText("Enter Species: ")
	record.CommonName, _ = userInput.GetText("Enter Common Name: ")
	record.StudySite, _ = userInput.GetText("Enter Study Site: ")
	record.AssociatedCommunity, _ = userInput.GetText("Enter Associated Community: ")
	record.Retinol, _ = userInput.GetText("Enter Retinol: ")

	return record
}

// AppendPreyRecord adds a new user-entered record to memory.
func AppendPreyRecord() {
	record := CreatePreyRecord()
	prey.Records = append(prey.Records, record)
	fmt.Print("Record created and appended.\n\n")
}

// EditPreyRecord replaces one existing record with newly entered values.
func EditPreyRecord(index int) {
	if index < 1 || index > len(prey.Records) {
		fmt.Printf("Invalid index. Please enter a number between 1 and %d.\n", len(prey.Records))
		return
	}

	DisplayPreyRecords(index)
	newRecord := CreatePreyRecord()
	prey.Records[index-1] = newRecord
	fmt.Print("Record updated.\n\n")
}

// DeletePreyRecord removes one prey record by its one-based index.
func DeletePreyRecord(index int) {
	if index < 1 || index > len(prey.Records) {
		fmt.Printf("Invalid index. Please enter a number between 1 and %d.\n", len(prey.Records))
		return
	}

	prey.Records = append(prey.Records[:index-1], prey.Records[index:]...)
	fmt.Print("Record deleted.\n\n")
}
