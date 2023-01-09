package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	var books [yearly]string
	books[0] = "The Hobbit"
	books[1] = "The Lord of the Rings"
	books[2] = "The Silmarillion"
	books[3] = books[0] + " and " + books[1]
	fmt.Printf("books : %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)
	wBooks[0] = books[0]
	// sBooks[0] = books[1]
	// sBooks[1] = books[2]
	// sBooks[2] = books[3]
	// or for i := 0; i < len(sBooks); i++ {
	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	fmt.Printf("\nwinter books : %#v\n", wBooks)
	fmt.Printf("\nsummer books : %#v\n", sBooks)

	var published [len(books)]bool
	published[0] = true
	published[len(books)-1] = true

	for i, pub := range published {
		if pub {
			fmt.Printf("Book %q is published\n", books[i])
		}
	}
}
