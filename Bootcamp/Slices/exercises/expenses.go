package main

import "fmt"

func main() {
	spendings := [][]int{
		{200, 100},
		{500},
		{50, 25, 75},
	}
	for i, daily := range spendings {
		// fmt.Println(i+1, daily)
		var total int
		for _, spending := range daily {
			total += spending
		}
		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}
