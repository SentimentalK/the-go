/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 21, 2026

References:
[1] The Go Authors. (n.d.). Package csv. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/encoding/csv [Accessed on: May 2026].

[2] The Go Authors. (n.d.). Package os. pkg.go.dev.
    [online]. Available at https://pkg.go.dev/os [Accessed on: May 2026].

[3] Fisheries and Oceans Canada. (2024, Dec. 16). Spatiotemporal variation in anadromous Arctic char (Salvelinus alpinus) foraging ecology and its influence on muscle pigmentation along western Hudson Bay, Nunavut, Canada. open.canada.ca.
    [online]. Available at https://open.canada.ca/data/en/dataset/9cbcf710-a2a1-11ef-8ccf-55cc7f028297 [Accessed on: Apr. 30, 2026].
*/

package model

// PreyRecord stores all values from one row of the prey CSV data.
type PreyRecord struct {
	Year                string
	Species             string
	CommonName          string
	StudySite           string
	AssociatedCommunity string
	Lat                 string
	Long                string
	Delta13C            string
	Delta13Cc           string
	Delta15N            string
	Delta15Nc           string
	TP                  string
	CN                  string
	AstaxanthinMgKg     string
	CanthaxanthinMgKg   string
	Retinol             string
}
