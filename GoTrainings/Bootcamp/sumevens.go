package main

import "fmt"

func main() {
	const (
		min = 1
		max = 10
	)
	var sum int
	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		sum = sum + i
		fmt.Print(i)
		if i < max {
			fmt.Print("+")
		}
	}
	fmt.Print("=", sum)
}
