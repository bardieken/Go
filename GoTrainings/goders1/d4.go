package main

import "fmt"

func main() {

	var score float64
	fmt.Println("enter Score")
	fmt.Scan(&score)
	switch {
	case score <= 59:
		fmt.Println("Yor grade is F")
	case score <= 69:
		fmt.Println("Your grade is c")
	default:
		fmt.Println("not known")

	}
}
