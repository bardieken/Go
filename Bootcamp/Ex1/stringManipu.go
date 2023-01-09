package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args
	if len(arg) != 3 {
		fmt.Println("Wrong command")
		return
	}
	command, word := arg[1], arg[2]
	switch command {
	case "lower":
		fmt.Println(strings.ToLower(word))
	case "upper":
		fmt.Println(strings.ToUpper(word))
	case "title":
		fmt.Println(strings.Title(word))
	default:
		fmt.Printf("Unknown command: %q\n", command)
	}
}
