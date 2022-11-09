package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("error:%q ins not a number", arg)
		return
	}
	m := feet * 0.3048
	fmt.Printf("%g feet is %g meters", feet, m)
}
