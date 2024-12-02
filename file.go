package elf

import (
	"bufio"
	"os"
)

// LinesFromFile reads lines from a file and returns them as a slice of strings.
// Panics if any error occurs during scanning.
func LinesFromFile(file *os.File) []string {
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return lines
}
