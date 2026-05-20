/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
*/

package main

import (
	"fmt"
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


	fmt.Println(filePath)
	fmt.Println(studentName, ":",studentID)
}