package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://www.merriam-webster.com/words-at-play/why-y-is-sometimes-a-vowel-usage
//  Check out: https://www.dictionary.com/e/w-vowel/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

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
