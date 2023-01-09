package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	// var books [yearly]string
	books := [...]string{
		"The Hobbit",
		"The Lord of the Rings",
		"The Silmarillion",
		"The Hobbit and The Lord of the Rings",
	}
	// fmt.Printf("books : %T\n", books)
	// fmt.Println("books : ", books)
	// fmt.Printf("books : %q\n", books)
	fmt.Printf("books : %#v\n", books)
}
