package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if a := os.Args; len(a) != 2 {
		fmt.Printf("%q is not a valid year\n", a[1])
		return
	}
	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year\n", os.Args[1])
		return
	}
	var leap bool
	if year%400 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	} else if year%4 == 0 {
		leap = true
	}
	if leap {
		fmt.Printf("%d is a leap year", year)
	} else {
		fmt.Printf("%d is not a leap year", year)
	}
}
