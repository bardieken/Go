package main

import "fmt"

func main() {
	var books [5]string // array
	books[0] = "draacula"
	books[1] = "1984"
	books[3] = "Island"
	//books = [5]string{
	//"dracula",
	//"1984",
	//"island",
	//}

	games := []string{"tuko", "puko"} // slice
	fmt.Printf("books\t: %#v\n", books)
	fmt.Printf("games\t: %#v\n", games)
}
