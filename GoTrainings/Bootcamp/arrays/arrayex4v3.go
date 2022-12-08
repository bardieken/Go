package main

import (
	"fmt"
	"strings"
)

func main() {
	books := []string{"The Hobbit", "The Lord of the Rings", "The Silmarillion"}
	upper := map(strings.ToUpper, books)
	lower := map(strings.ToLower, books)

	fmt.Printf("books: %q\n", books)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)
}
