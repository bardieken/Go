package main

import "fmt"

func main() {
	word := "numberofcharacters"
	var i int
	for range word {
		i++
	}

	fmt.Printf("There are %d characters in word\n", i)
}
