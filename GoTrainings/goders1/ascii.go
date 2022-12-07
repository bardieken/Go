package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Get the text to convert from the command line arguments.
	text := strings.Split(os.Args[1], "\\n")

	// Read the font file.
	fontBytes, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		panic(err)
	}

	// Split the font file into lines.
	fontLines := strings.Split(string(fontBytes), "\n")

	// Get the maximum length of the font lines.
	maxLineLength := 0
	for _, line := range fontLines {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}

	// Create a 2D slice to store the ASCII art.
	art := make([][]rune, len(fontLines))
	for i := range art {
		art[i] = make([]rune, maxLineLength)
	}

	// Loop through the text and the font lines, adding each character
	// from the text to the corresponding character in the font line.
	for i, line := range fontLines {
		for j, char := range line {
			if j < len(text) {
				art[i][j] = rune([]byte(text[j])[0])
			} else {
				art[i][j] = char
			}
		}
	}

	// Print the ASCII art to the terminal.
	for _, line := range art {
		fmt.Println(string(line))
	}
}
