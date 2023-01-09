package main

import "fmt"

func main() {
	prev := [3]string{"Hobbits", "Hobbits time", "Hobbits end"}
	books := prev
	var adbook [4]string

	for i, b := range prev {
		books[i] += " 2nd edition"
		adbook[i] += b + " 2nd edition"
	}
	adbook[3] = "Hobbits reborn"
	fmt.Printf("Books: %q\n", prev)
	fmt.Printf("new Books: %q\n", books)
	fmt.Printf("Additional Books: %q\n", adbook)
}
