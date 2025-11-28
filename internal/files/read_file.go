// Package files provides utilities for file operations.
package files

import (
	"bufio"
	"os"
)

// ReadFile reads a file line by line and returns a slice of strings.
func ReadFile(filename string) (lines []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(wd + "/" + filename)
	if err != nil {
		panic(err)
	}

	buffer := bufio.NewScanner(file)

	for buffer.Scan() {
		lines = append(lines, buffer.Text())
	}

	file.Close()

	return lines
}

// ReadFileStr reads the entire content of a file and returns it as a string.
func ReadFileStr(filename string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(wd + "/" + filename)
	if err != nil {
		panic(err)
	}

	return string(data)
}
