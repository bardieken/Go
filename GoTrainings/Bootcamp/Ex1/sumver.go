package main

import "fmt"

func main() {
	const min, max int = 1, 10
	var sum int
	for i := min; i <= max; i++ {
		sum += i
		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" =%d\n", sum)
}
