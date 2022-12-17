package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EUR = iota
	GBP
	JPY
)

func main() {
	args := os.Args
	if len(args) < 1 {
		fmt.Println(" Please provide the amount to be converted.")
		os.Exit(0)
	}
	rates := [...]float64{0.88, 0.78, 113.08}

	a, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
	}
	fmt.Printf("%.2f USD = %.2f EUR\n", a, rates[EUR]*a)
	fmt.Printf("%.2f USD = %.2f GBP\n", a, rates[GBP]*a)
	fmt.Printf("%.2f USD = %.2f JPY\n", a, rates[JPY]*a)
}
