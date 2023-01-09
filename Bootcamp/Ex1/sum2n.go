package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
	func main() {
		args := os.Args
		_, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(" Only numbers are allowed")
		}

		min, err := strconv.Atoi(args[1])
		max, err := strconv.Atoi(args[2])
		if len(args) != 3 {
			fmt.Println(" Usage: wite min and max")
			os.Exit(1)
		} else if min > max {
			fmt.Println("min < max")
			os.Exit(1)
		}
		var sum int
		for i := min; i <= max; i++ {
			sum += i
			fmt.Print(i)
			if i < max {
				fmt.Print("+")
			}
		}
		fmt.Println("=", sum)
	}
*/
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: min max")
		return
	}
	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || min > max {
		fmt.Println("wrong input")
		return
	}
	var sum int
	for i := min; i <= max; i++ {
		sum = sum + i
		fmt.Print(i)
		if i < max {
			fmt.Print("+")
		}
	}
	fmt.Println("=", sum)
}
