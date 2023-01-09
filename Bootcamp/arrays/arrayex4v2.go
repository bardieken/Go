package main

import (
	"fmt"
	"strings"
)

func main() {
	books := [3]string{"The Hobbit", "The Lord of the Rings", "The Silmarillion"}
	// create new arrays
	upper := make([]string, len(books))
	lower := make([]string, len(books))
	// irritate over the books array
	for i, book := range books {
		upper[i] = strings.ToUpper(book)
		lower[i] = strings.ToLower(book)
	}
	fmt.Printf("books: %q\n", books)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)
}
