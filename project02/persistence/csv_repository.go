/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 21, 2026

References:
[1] The Go Authors. (n.d.). Package csv. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/encoding/csv [Accessed on: May 2026].

[2] The Go Authors. (n.d.). Package fmt. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/fmt [Accessed on: May 2026].

[3] The Go Authors. (n.d.). Package os. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/os [Accessed on: May 2026].

[4] Google. (n.d.). Package uuid. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/github.com/google/uuid [Accessed on: May 2026].

[5] Fisheries and Oceans Canada. (2024, Dec. 16). Spatiotemporal variation in anadromous Arctic char (Salvelinus alpinus) foraging ecology and its influence on muscle pigmentation along western Hudson Bay, Nunavut, Canada. open.canada.ca.
    [online]. Available at https://open.canada.ca/data/en/dataset/9cbcf710-a2a1-11ef-8ccf-55cc7f028297 [Accessed on: Apr. 30, 2026].
*/

package persistence

import (
	"encoding/csv"
	"fmt"
	"os"
	"project02/model"

	"github.com/google/uuid"
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
		if len(row) < 16 {
			return nil, fmt.Errorf("row %d has %d columns; expected at least 16", index+1, len(row))
		}
		record := model.PreyRecord{
			Year:                row[0],
			Species:             row[1],
			CommonName:          row[2],
			StudySite:           row[3],
			AssociatedCommunity: row[4],
			Lat:                 row[5],
			Long:                row[6],
			Delta13C:            row[7],
			Delta13Cc:           row[8],
			Delta15N:            row[9],
			Delta15Nc:           row[10],
			TP:                  row[11],
			CN:                  row[12],
			AstaxanthinMgKg:     row[13],
			CanthaxanthinMgKg:   row[14],
			Retinol:             row[15],
		}
		records = append(records, record)
		if index >= limit {
			break
		}
	}

	return records, nil
}

// WriteToCSV writes prey records to a new CSV file with a unique file name.
func WriteToCSV(records []model.PreyRecord) error {
	// Create or overwrite the CSV file
	filePath := fmt.Sprintf("data/prey_records_%s.csv", uuid.NewString())
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Year", "Species", "Common Name", "Study Site", "Associated Community", "Lat", "Long", "Delta13C", "Delta13Cc", "Delta15N", "Delta15Nc", "TP", "C:N", "Astaxanthin", "Canthaxanthin", "Retinol"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write records
	for _, record := range records {
		row := []string{
			record.Year,
			record.Species,
			record.CommonName,
			record.StudySite,
			record.AssociatedCommunity,
			record.Lat,
			record.Long,
			record.Delta13C,
			record.Delta13Cc,
			record.Delta15N,
			record.Delta15Nc,
			record.TP,
			record.CN,
			record.AstaxanthinMgKg,
			record.CanthaxanthinMgKg,
			record.Retinol,
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	fmt.Printf("Records saved to %s\n\n", filePath)
	return nil
}
