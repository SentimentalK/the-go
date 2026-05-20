/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda

References:
[1] The Go Authors, "Package csv," pkg.go.dev. [Online]. Available: https://pkg.go.dev/encoding/csv. [Accessed: May 2026].
[2] The Go Authors, "Package os," pkg.go.dev. [Online]. Available: https://pkg.go.dev/os. [Accessed: May 2026].
[3] Fisheries and Oceans Canada, "Spatiotemporal variation in anadromous Arctic char..." open.canada.ca. [Online]. Available: https://open.canada.ca/data/en/dataset/9cbcf710-a2a1-11ef-8ccf-55cc7f028297. [Accessed: Apr. 30, 2026].
*/

package model

// PreyRecord stores selected values from one row of the prey CSV data.
type PreyRecord struct {
	Year                string
	Species             string
	CommonName          string
	StudySite           string
	AssociatedCommunity string
	Retinol             string
}
