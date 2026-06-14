/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 13, 2026

References:
[1] The Go Authors. (n.d.). Package testing. pkg.go.dev.
    [online]. Available at: https://pkg.go.dev/testing [Accessed on: June 14, 2026].
*/

// Package calc provides calculator operations and a console menu.
package calc

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers.
func Subtract(a, b int) int {
	return a - b
}
