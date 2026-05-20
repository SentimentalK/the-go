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
	"encoding/csv"
	"fmt"
	"os"
	"project01/model"
)

func main() {
	// Automatically detect type of variable based on assigned value
	studentName := "Xinghan Xu"
	// Explicitly declare type of variable
	var filePath string = "./data/Prey collection & analysis - raw data.csv"
	// Constrants
	const studentID = "041164952"
	records := []model.PreyRecord{}

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

	limit := 5
	for index, row := range rows {
		if index < 1 {
			continue
		}
		record := model.PreyRecord{
			Year:                row[0],
			Species:             row[1],
			CommonName:          row[2],
			StudySite:           row[3],
			AssociatedCommunity: row[4],
			Retinol:             row[5],
		}
		records = append(records, record)
		if index >= limit {
			break
		}
	}

	fmt.Println(studentName, ":", studentID)
	fmt.Println("Total records read:", len(records))
	fmt.Println("---")
	for _, record := range records {
		fmt.Println("Year:", record.Year)
		fmt.Println("Species:", record.Species)
		fmt.Println("Common Name:", record.CommonName)
		fmt.Println("Study Site:", record.StudySite)
		fmt.Println("Associated Community:", record.AssociatedCommunity)
		fmt.Println("Retinol:", record.Retinol)
		fmt.Println("---")
	}

}
