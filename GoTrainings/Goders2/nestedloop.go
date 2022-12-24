package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("For Outer loop : %d \t Inner loop is : %d\n", i, j)
		}
	}
}
