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

	// fmt.Printf("books : %T\n", books)
	// fmt.Println("books : ", books)
	// fmt.Printf("books : %q\n", books)
	fmt.Printf("books : %#v\n", books)
}
