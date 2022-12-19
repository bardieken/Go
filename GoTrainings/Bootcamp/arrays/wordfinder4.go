package main

import (
	"fmt"
	"os"
	"strings"
)

var filter = map[string]bool{
	"and":   true,
	"or":    true,
	"was":   true,
	"the":   true,
	"since": true,
	"very":  true,
}

const text = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	for _, arg := range os.Args[1:] {
		if len(arg) < 1 {
			fmt.Println("Please give me a word to search.")
			os.Exit(0)
		}
		title := strings.ToLower(arg)
		if filter[title] {
			fmt.Println("")
			continue
		}

		for i, input := range strings.Fields(text) {
			search := strings.Contains(strings.ToLower(input), title)
			if search {
				fmt.Printf("%#v : %q\n", i+1, input)
			}
		}
	}
}
