package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Books := []string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Tell me a book title")
		os.Exit(0)
	}
	title := strings.ToLower(args[1])
	for _, book := range Books {
		search := strings.Contains(strings.ToLower(book), title)
		if search {
			fmt.Printf("Search Results:\n + %s\n", book)
		}
	}
}
