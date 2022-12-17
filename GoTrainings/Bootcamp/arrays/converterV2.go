package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const (
		USD = 1.0
		EUR = USD * 0.88
		GBP = USD * 0.78
		JPY = USD * 113.08
	)
	args := os.Args
	if len(args) < 2 {
		fmt.Println(" Please provide the amount to be converted.")
		fmt.Println(len(args))
		os.Exit(0)
	}
	a, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
	}
	fmt.Printf("%.2f USD = %.2f EUR\n", a, EUR*a)
	fmt.Printf("%.2f USD = %.2f GBP\n", a, GBP*a)
	fmt.Printf("%.2f USD = %.2f JPY\n", a, JPY*a)
}
