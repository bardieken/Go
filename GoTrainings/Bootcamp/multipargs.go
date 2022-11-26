package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	number := os.Args[1]
	if len(os.Args) < 1 {
		fmt.Println("Give a number")
		return
	}
	num, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("not a number")
		return
	}
	fmt.Printf("%5s", "X")
	for i := 0; i <= num; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	for i := 0; i <= num; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= num; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
