/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 21, 2026

References:
[1] The Go Authors. (n.d.). Package testing. pkg.go.dev.
    [Online]. Available at https://pkg.go.dev/testing [Accessed on: June 15, 2026].
*/

package tests

import (
	"project03/persistence"
	"testing"
)

// TestReadLinesForLoadsNoMoreThanOneHundredRecords checks the dataset load limit.
func TestReadLinesForLoadsNoMoreThanOneHundredRecords(t *testing.T) {
	records, err := persistence.ReadLinesFor(100, "../data/Prey collection & analysis - raw data.csv")
	if err != nil {
		t.Fatalf("ReadLinesFor returned error: %v", err)
	}

	if len(records) == 0 {
		t.Fatal("expected records to be loaded")
	}
	if len(records) > 100 {
		t.Fatalf("expected at most 100 records, got %d", len(records))
	}
	if records[0].Year == "" || records[0].Species == "" || records[0].Retinol == "" {
		t.Fatalf("expected first record fields to be populated, got %+v", records[0])
	}
}
