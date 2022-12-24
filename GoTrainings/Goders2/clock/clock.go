package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("clock.txt")
	if err != nil {
		os.Exit(0)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		os.Exit(0)
	}
	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines); i += 5 {
		for j := i; j < i+5; j++ {
			fmt.Printf("%s  ", lines[j])
		}
		fmt.Println()
	}
}
