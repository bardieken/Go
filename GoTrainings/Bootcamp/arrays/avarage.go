package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 || len(args) > 6 {
		fmt.Println("Please tell me numbers (maximum 5 numbers)")
		// fmt.Println(len(args))
		os.Exit(0)
	}
	for 
		
	}

	numbers := args
	var avrg float64
	for _, number := range numbers {
		avrg += float64(number)
	}
	fmt.Println("Your numbers are: ", numbers)
	fmt.Println("Avarage: ")
}
