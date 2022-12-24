package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if a := os.Args; len(a) == 1 {
		fmt.Println("Set Age")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("%q is not a valid age\n", a[1])
	} else if n < 0 {
		fmt.Println("Age can not be a negative number")
	} else if n < 13 {
		fmt.Println("PG-Rated")
	} else if n > 13 && n < 17 {
		fmt.Println("PG-13")
	} else {
		fmt.Println("R-Rated")
	}
}
