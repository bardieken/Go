package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	capsOn := false
	for _, arg := range arguments {
		if arg == "--upper" {
			capsOn = true
		}
	}
	for _, arg := range arguments {
		numv := 0
		for _, v := range arg {
			numv = numv*10 + int(v-'0')
		}
		if numv >= 1 && numv <= 26 {
			if capsOn == false {
				z01.PrintRune(rune(numv + 96))
			} else {
				z01.PrintRune(rune(numv + 64))
			}
		} else {
			if arg != "--upper" {
				z01.PrintRune(' ')
			}
		}
	}
	if len(arguments) > 0 {
		z01.PrintRune('\n')
	}
}
