/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 13, 2026

References:
[1] The Go Authors. (n.d.). Package testing. pkg.go.dev.
    [online]. Available at: https://pkg.go.dev/testing [Accessed on: June 14, 2026].
*/

package calc

import "testing"

// TestAdd verifies that Add returns the sum of two integers.
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// TestSubtract verifies that Subtract returns the difference between two integers.
func TestSubtract(t *testing.T) {
	result := Subtract(5, 2)
	expected := 4
	if result != expected {
		t.Errorf("Subtract(5, 2) = %d; want %d", result, expected)
	}
}
