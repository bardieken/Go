package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(" Magnitude missing")
		return
	}
	magni, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println(" Magnitude missing")
		return
	}
	switch {
	case magni < 2.0:
		fmt.Println("micro")
	case magni >= 2.0 && magni < 3.0:
		fmt.Println("very minor")
	case magni >= 3.0 && magni < 4.0:
		fmt.Println("minor")
	case magni >= 4.0 && magni < 5.0:
		fmt.Println("light")
	case magni >= 5.0 && magni < 6.0:
		fmt.Println("moderate")
	case magni >= 6.0 && magni < 7.0:
		fmt.Println("strong")
	case magni >= 7.0 && magni < 8.0:
		fmt.Println("major")
	case magni >= 8.0 && magni < 10.0:
		fmt.Println("great")
	case magni >= 10.0:
		fmt.Println("massive")
	default:
		fmt.Println("Magnitude missing")
	}
}
