package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}
	var richter string
	// desc := args[1]
	switch args[1] {
	case "micro":
		richter = "Less than 2.0"
	case "very minor":
		richter = "2.0 to 2.9"
	case "minor":
		richter = "3.0 to 3.9"
	case "light":
		richter = "4.0 to 4.9"
	case "moderate":
		richter = "5.0 to 5.9"
	case "strong":
		richter = "6.0 to 6.9"

	case "major":
		richter = "7.0 to 7.9"
	case "great":
		richter = "8.0 to 9.9"
	case "massive":
		richter = "10.0 or greater"
	default:
		richter = "Magnitude missing"
	}
	fmt.Printf("%s's richter  is %s\n", args[1], richter)
}
