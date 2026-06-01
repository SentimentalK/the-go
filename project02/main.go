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
	"project02/presentation"
)

// main reads a fixed number of prey records from the CSV file and prints them.
func main() {
	presentation.MenuLoop()
}
