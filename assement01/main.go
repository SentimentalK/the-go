/*
Author: Xinghan Xu
Course: CST8002 Programming Language Research Project

References:
[1] The Go Authors, “Tutorial: Get started with Go,” The Go Programming Language.
    [Online]. Available: https://go.dev/doc/tutorial/getting-started.
    [Accessed: May 14, 2026].

[2] The Go Authors, “Go Doc Comments,” The Go Programming Language.
    [Online]. Available: https://go.dev/doc/comment.
    [Accessed: May 14, 2026].

[3] The Go Authors, “Effective Go,” The Go Programming Language.
    [Online]. Available: https://go.dev/doc/effective_go.
    [Accessed: May 14, 2026].

[4] The Go Authors, “Command Documentation,” The Go Programming Language.
    [Online]. Available: https://go.dev/doc/cmd.
    [Accessed: May 14, 2026].
*/

// Package main demonstrates a small console program written in Go.
package main

import (
	"fmt"
	"runtime"
)

const studentName = "Xinghan Xu"

// main is the entry point of the program.
func main() {
	printProgramInformation()
}

// printProgramInformation prints the required proof-of-concept output.
func printProgramInformation() {
	fmt.Println("Hello Tuna Fish!")
	fmt.Printf("Programming language: Go, version: %s\n", runtime.Version())
	fmt.Printf("Student name: %s\n", studentName)
}