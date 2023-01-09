package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 || len(args) > 5 {
		fmt.Println("Plese write up to 5 numbers")
		return
	}
	var sum float64
	var total float64

	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			continue
		}
		sum += number
		total++
	}
	avarage := sum / total
	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("Average: %.2f\n", avarage)
}
