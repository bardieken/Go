package main

import "fmt"

func main() {
	names := [3]string{"Einstein", "Newton", "Copernicus"}
	books := [5]string{"The Hobbit", "The Lord of the Rings", "The Silmarillion"}
	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}
