package business

import (
	"fmt"
	"project02/model"
	"project02/persistence"
	"project02/presentation/user"
)

type Prey struct {
	Records []model.PreyRecord
}

var prey = &Prey{}
var userInput = user.NewConsoleInput()

func LoadPreyRecords(filePath string) ([]model.PreyRecord, error) {
	records, err := persistence.ReadLinesFor(100, filePath)
	if err != nil {
		return nil, err
	}

	prey.Records = records
	fmt.Printf("%d records loaded.\n", len(records))
	return records, nil
}

func SavePreyRecords() error {
	if prey.Records == nil {
		return fmt.Errorf("no records to save")
	}

	return persistence.WriteToCSV(prey.Records)
}

func DisplayPreyRecords(index int) {

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

func AppendPreyRecord() {
	record := CreatePreyRecord()
	prey.Records = append(prey.Records, record)
}

func EditPreyRecord(index int) {
	if index < 1 || index > len(prey.Records) {
		fmt.Printf("Invalid index. Please enter a number between 1 and %d.\n", len(prey.Records))
		return
	}

	DisplayPreyRecords(index)
	newRecord := CreatePreyRecord()
	prey.Records[index-1] = newRecord
}

func DeletePreyRecord(index int) {
	if index < 1 || index > len(prey.Records) {
		fmt.Printf("Invalid index. Please enter a number between 1 and %d.\n", len(prey.Records))
		return
	}

	prey.Records = append(prey.Records[:index-1], prey.Records[index:]...)
}
