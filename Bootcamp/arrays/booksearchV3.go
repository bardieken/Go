package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := [4]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Tell me a book title")
		return
	}
	var found bool
	search := strings.ToLower(args[0])
	fmt.Println("Search Results:")
	for _, book := range books {
		words := strings.Fields(book)
		for _, word := range words {
			if strings.EqualFold(word, search) {
				fmt.Printf(" + %s\n", book)
				found = true
				break
			}
		}
	}
	if !found {
		fmt.Println("No results found")
	}
}
