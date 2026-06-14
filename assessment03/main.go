/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project
Professor: Stanley Pieda
Due Date: June 13, 2026

References:
[1] The Go Authors. (n.d.). Command go. go.dev.
    [online]. Available at: https://go.dev/cmd/go/ [Accessed on: June 14, 2026].
*/

// Package main starts the Assessment 03 calculator program.
package main

import (
	"assessment03/calc"
	"fmt"
)

// main is the entry point of the program.
func main() {
	fmt.Println("Program by Xinghan Xu")
	fmt.Println(calc.Add(2, 3))
	fmt.Println(calc.Subtract(5, 2))
}
