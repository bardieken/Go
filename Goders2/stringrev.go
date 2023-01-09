package main

import (
	"fmt"
)

/*
	func main() {
		word := "numberofcharacters"
		var a []string
		for _, r := range word {
			a = append(a, string(r))
		}
		for i := len(a) - 1; i >= 0; i-- {
			fmt.Print(a[i])
		}
		fmt.Println()
	}
*/
func main() {
	word := "numberofcharacters"
	for i := len(word) - 1; i >= 0; i-- {
		fmt.Print(string(word[i]))
	}
	fmt.Println()
}
