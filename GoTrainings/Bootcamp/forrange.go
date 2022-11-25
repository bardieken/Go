package main

import (
	"fmt"
	"strings"
)

/*
	func main() {
		for _, v := range os.Args[1:] {
			fmt.Printf("%q\n", v)
		}
	}
*/
func main() {
	words := strings.Fields(
		" birinci ikinci ucuncu dorduncu besinci altinci",
	)
	var (
		i int
		v string // boylece range degerleri bir sonraki islem icin korur
	)
	// for i, v = range words yerine asagidaki gibi de yazilabilir =
	for i, v = range words {
		fmt.Printf("%-2d: %q\n", i+1, v)
	}
}

// range len(words) kadar doner ve her donuste i ve v yi gunceller.
// i ve v nin degerleri bir sonraki islem icin korunur.
// range returns next index and value
