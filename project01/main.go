/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda

References:
[1] The Go Authors, "Package csv," pkg.go.dev. [Online]. Available: https://pkg.go.dev/encoding/csv. [Accessed: May 2026].
[2] The Go Authors, "Package os," pkg.go.dev. [Online]. Available: https://pkg.go.dev/os. [Accessed: May 2026].
[3] Fisheries and Oceans Canada, "Spatiotemporal variation in anadromous Arctic char..." open.canada.ca. [Online]. Available: https://open.canada.ca/data/en/dataset/9cbcf710-a2a1-11ef-8ccf-55cc7f028297. [Accessed: Apr. 30, 2026].
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
	for index, row := range rows {
		if index < 1 {
			continue
		}
		record := PreyRecord{
			Year:                row[0],
			Species:             row[1],
			CommonName:          row[2],
			StudySite:           row[3],
			AssociatedCommunity: row[4],
			Retinol:             row[5],
		}
		records = append(records, record)
		break
	}

	fmt.Println(studentName, ":",studentID)
	fmt.Println("Total records read:", len(records)-1)
	fmt.Println("Year:", records[1].Year)
	fmt.Println("Species:", records[1].Species)
	fmt.Println("Common Name:", records[1].CommonName)
	fmt.Println("Study Site:", records[1].StudySite)
	fmt.Println("Associated Community:", records[1].AssociatedCommunity)
	fmt.Println("Retinol:", records[1].Retinol)

}