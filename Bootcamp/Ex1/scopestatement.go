package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var n int
	if a := os.Args; len(a) == 1 {
		fmt.Println(" Give Number")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("%q is not a valid number\n", a[1])
	} else {
		fmt.Printf("%s * 2 = %d\n", a[1], n*2)
	}
	// fmt.Println(n)
}
