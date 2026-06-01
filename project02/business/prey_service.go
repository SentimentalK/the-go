package business

import (
	"fmt"
	"bufio"
	"strings"
	"project02/model"
	"project02/persistence"
)

func (prey *Prey) LoadPreyRecords(filePath string) ([]model.PreyRecord, error) {
	return persistence.ReadLinesFor(100, filePath)
}

func (prey *Prey) SavePreyRecords() error {
	return persistence.WriteToCSV(prey.Records)
}

func (prey *Prey) DisplayPreyRecords(index int) {
	if index < 1 || index > len(prey.Records) {
		fmt.Printf("Invalid index. Please enter a number between 1 and %d.\n", len(prey.Records))
		return
	}
	record := prey.Records[index-1]
	fmt.Printf("Record %d:\n", index)
	fmt.Printf("Year: %s\n", record.Year)
	fmt.Printf("Species: %s\n", record.Species)
	fmt.Printf("Common Name: %s\n", record.CommonName)
	fmt.Printf("Study Site: %s\n", record.StudySite)
	fmt.Printf("Associated Community: %s\n", record.AssociatedCommunity)
	fmt.Printf("Retinol: %s\n", record.Retinol)
}

func (prey *Prey, userInput *presentation.ConsoleInput) CreatePreyRecord() model.PreyRecord {
	var record model.PreyRecord
	record.Year = userInput.GetText("Enter Year: ")
	record.Species = userInput.GetText("Enter Species: ")
	record.CommonName = userInput.GetText("Enter Common Name: ")
	record.StudySite = userInput.GetText("Enter Study Site: ")
	record.AssociatedCommunity = userInput.GetText("Enter Associated Community: ")
	record.Retinol = userInput.GetText("Enter Retinol: ")

	return record
}

func (prey *Prey) AppendPreyRecord() {
	record := CreatePreyRecord(userInput)
	prey.Records = append(prey.Records, record)
}

func (prey *Prey) EditPreyRecord(index int) {
	DisplayPreyRecords(index)
	newRecord := CreatePreyRecord()
	prey.Records[index-1] = newRecord
}

func (prey *Prey) DeletePreyRecord(index int) {
	if index < 1 || index > len(prey.Records) {
		fmt.Printf("Invalid index. Please enter a number between 1 and %d.\n", len(prey.Records))
		return
	}
	prey.Records = append(prey.Records[:index-1], prey.Records[index:]...)
}
