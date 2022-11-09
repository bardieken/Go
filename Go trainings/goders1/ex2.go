package main

import "fmt"

func main() {
	var gun int
	fmt.Println("Kacinci gun?")
	fmt.Scan(&gun)
	switch gun {
	case 1:
		fmt.Println("Pazartesi")
	case 2:
		fmt.Println("sali")
	default:
		fmt.Println("sie")
	}
}
