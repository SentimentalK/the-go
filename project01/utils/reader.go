/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda

References:
[1] The Go Authors, "Package csv," pkg.go.dev. [Online]. Available: https://pkg.go.dev/encoding/csv. [Accessed: May 2026].
[2] The Go Authors, "Package os," pkg.go.dev. [Online]. Available: https://pkg.go.dev/os. [Accessed: May 2026].
[3] Fisheries and Oceans Canada, "Spatiotemporal variation in anadromous Arctic char..." open.canada.ca. [Online]. Available: https://open.canada.ca/data/en/dataset/9cbcf710-a2a1-11ef-8ccf-55cc7f028297. [Accessed: Apr. 30, 2026].
*/

package utils

import (
	"encoding/csv"
	"os"
	"project01/model"
)

// ReadLinesFor reads up to n prey records from the CSV file at filePath.
func ReadLinesFor(n int, filePath string) ([]model.PreyRecord, error) {

	records := []model.PreyRecord{}

	// Open File
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read csv file
	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	limit := n
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

	return records, nil
}
