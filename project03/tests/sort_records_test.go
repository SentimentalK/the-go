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
	"project03/business"
	"testing"
)

func TestSortPreyRecordsByLat(t *testing.T) {
	// load records
	records, err := business.LoadPreyRecords("../data/Prey collection & analysis - raw data.csv")
	if err != nil {
		t.Fatalf("Error loading records: %v", err)
	}

	// call sort
	business.SortPreyRecordsByLat()

	// check Lat order
	for i := 0; i < len(records)-1; i++ {
		if records[i].Lat > records[i+1].Lat {
			t.Errorf("Records are not sorted by Lat. Index %d has Lat %s, Index %d has Lat %s", i, records[i].Lat, i+1, records[i+1].Lat)
		}
	}
}
