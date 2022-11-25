package main

import "fmt"

func main() {
	number := 5
	fmt.Printf("%5s", "X")
	for i := 0; i <= number; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	for i := 0; i <= number; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= number; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
