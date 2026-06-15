/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 21, 2026

References:
[1] The Go Authors. (n.d.). Package fmt. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/fmt [Accessed on: June 1, 2026].

[2] The Go Authors. (n.d.). Package reflect. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/reflect [Accessed on: June 1, 2026].

[3] Fisheries and Oceans Canada. (2024, Dec. 16). Spatiotemporal variation in anadromous Arctic char (Salvelinus alpinus) foraging ecology and its influence on muscle pigmentation along western Hudson Bay, Nunavut, Canada. open.canada.ca.
    [online]. Available at https://open.canada.ca/data/en/dataset/9cbcf710-a2a1-11ef-8ccf-55cc7f028297 [Accessed on: Apr. 30, 2026].
*/

package business

import (
	"fmt"
	"project03/model"
	"project03/persistence"
	"project03/presentation/user"
	"reflect"
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
	recordValue := reflect.ValueOf(record)
	recordType := recordValue.Type()

	fmt.Printf("\n--Record %d:--\n", index)
	for fieldIndex := 0; fieldIndex < recordValue.NumField(); fieldIndex++ {
		fmt.Printf("%s: %s\n", recordType.Field(fieldIndex).Name, recordValue.Field(fieldIndex).String())
	}
	fmt.Print("---\n\n")
}

// CreatePreyRecord asks the user for record values and returns a new prey record.
func CreatePreyRecord() model.PreyRecord {
	var record model.PreyRecord
	recordValue := reflect.ValueOf(&record).Elem()
	recordType := recordValue.Type()

	for fieldIndex := 0; fieldIndex < recordValue.NumField(); fieldIndex++ {
		text, _ := userInput.GetText(fmt.Sprintf("Enter %s: ", recordType.Field(fieldIndex).Name))
		recordValue.Field(fieldIndex).SetString(text)
	}

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
