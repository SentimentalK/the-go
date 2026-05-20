/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
*/

package main

import (
	"fmt"
	"os"
	"encoding/csv"
)

func main() {
	// Automatically detect type of variable based on assigned value
	studentName := "Xinghan Xu"
	// Explicitly declare type of variable
	var filePath string = "./data/Prey collection & analysis - raw data.csv"
	// Constrants
	const studentID = "041164952"

	// Go does not have classes or "objects" as defined by languages like Java
	// A struct is strictly a composite data type—a blueprint for a block of memory 
	// containing a collection of named fields.
	type PreyRecord struct {
		Year                string
		Species             string
		CommonName          string
		StudySite           string
		AssociatedCommunity string
		Retinol             string
	}
	// Create an instance of the struct
	record := PreyRecord{
		Year:                "2024",
		Species:             "Salvelinus alpinus",
		CommonName:          "Arctic char",
		StudySite:           "Example Site",
		AssociatedCommunity: "Example Community",
		Retinol:             "12.5",
	}
	// slice
	records := []PreyRecord{}
	records = append(records, record)
	// Print in loop
	for index, record := range records {
		fmt.Println("Record", index+1)
		fmt.Println("Year:", record.Year)
		fmt.Println("Species:", record.Species)
	}

	// Open File
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// read csv file
	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Print the first row
	for _, row := range rows {
		fmt.Println(row)
		break
	}

	fmt.Println(studentName, ":",studentID)
}