package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if a := os.Args; len(a) != 2 {
		fmt.Println("Give a number")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("%q id not a valid number", a[1])
	} else if n%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8", n)
	} else if n%2 == 0 && n%8 != 0 {
		fmt.Printf("%d is an even number", n)
	} else {
		fmt.Printf("%d is an odd number", n)
	}
}
