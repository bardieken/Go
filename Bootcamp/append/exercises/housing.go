package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		location []string // sizes    []int
		size     []string
		beds     []string
		baths    []string
		prices   []string
	)
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)
	head := strings.Split(header, ",")
	dat := strings.Split(data, "\n")
	for i := range dat {
		da := strings.Split(dat[i], ",")
		location = append(location, da[0])

		size = append(size, da[1])

		beds = append(beds, da[2])

		baths = append(baths, da[3])

		prices = append(prices, da[4])
	}
	fmt.Printf("%-15s\n", head)
	i := 0
	for i < 80 {
		i++
		fmt.Printf("=")
	}
	fmt.Println()
	for i := range location {
		fmt.Printf("%-15s %-15s %-15s %-15s %-15s\n", location[i], size[i], beds[i], baths[i], prices[i])
	}
}
