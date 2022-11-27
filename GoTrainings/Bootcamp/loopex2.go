package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}
	op := args[1]
	size, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Size is not a number")
		return
	}
	switch op {
	case "*":
		if size < 0 {
			fmt.Println("Size must be positive")
			return
		}
		fmt.Printf("%5s", "X")
		for i := 0; i <= size; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()
		for i := 0; i <= size; i++ {
			fmt.Printf("%5d", i)
			for j := 0; j <= size; j++ {
				fmt.Printf("%5d", i*j)
			}
			fmt.Println()
		}
	case "/":
		if size < 0 {
			fmt.Println("Size must be positive")
			return
		}
		fmt.Printf("%5s", "X")
		for i := 1; i <= size; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()
		for i := 1; i <= size; i++ {
			fmt.Printf("%5d", i)
			for j := 1; j <= size; j++ {
				fmt.Printf("%5d", i/j)
			}
			fmt.Println()
		}
	}
}
