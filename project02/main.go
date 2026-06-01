/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda

References:
[1] The Go Authors. (n.d.). Package csv. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/encoding/csv [Accessed on: May 2026].

[2] The Go Authors. (n.d.). Package os. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/os [Accessed on: May 2026].

[3] Fisheries and Oceans Canada. (2024, Dec. 16). Spatiotemporal variation in anadromous Arctic char (Salvelinus alpinus) foraging ecology and its influence on muscle pigmentation along western Hudson Bay, Nunavut, Canada. open.canada.ca.
    [online]. Available at https://open.canada.ca/data/en/dataset/9cbcf710-a2a1-11ef-8ccf-55cc7f028297 [Accessed on: Apr. 30, 2026].
*/

package main

import (
	"fmt"
	"project02/persistence"
)

// main reads a fixed number of prey records from the CSV file and prints them.
func main() {
	// Automatically detect type of variable based on assigned value
	studentName := "Xinghan Xu"
	// Explicitly declare type of variable
	var filePath string = "./data/Prey collection & analysis - raw data.csv"
	// Constants
	const studentID = "041164952"

	records, err := persistence.ReadLinesFor(5, filePath)
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
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
