package main

import (
	"fmt"
	"strings"
)

/*
	func main() {
		for i := 1; i < len(os.Args); i++ {
			fmt.Println(os.Args[i])
		}
	}
*/
func main() {
	words := strings.Fields(
		" birinci ikinci ucuncu dorduncu besinci altinci",
	)
	for j := 0; j < len(words); j++ {
		fmt.Printf("%-2d: %q\n", j+1, words[j])
	}
}
