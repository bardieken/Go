package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]
	l := len(os.Args[1])
	if l > 1 || l == 0 {
		fmt.Println("Give me a letter")
	} else if arg == "y" || arg == "w" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.", arg)
	} else if arg == "a" || arg == "e" || arg == "i" || arg == "o" || arg == "u" {
		fmt.Printf("%q is a vowel.\n", arg)
	} else {
		fmt.Printf("%q is a consonant.\n", arg)
	}
}
