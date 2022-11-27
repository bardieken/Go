package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num := os.Args[1]
	if len(os.Args) != 2 {
		fmt.Println("Please enter a number")
		os.Exit(1)
	} else if n, err := strconv.Atoi(num); err != nil || n < 0 {
		fmt.Printf("%q is not a valid number\n", num)
		os.Exit(1)
	} else {
		fmt.Printf("%5s", "X")
		for i := 0; i <= n; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()
		for i := 0; i <= n; i++ {
			fmt.Printf("%5d", i)
			for j := 0; j <= n; j++ {
				fmt.Printf("%5d", i*j)
			}
			fmt.Println()
		}
	}
}
