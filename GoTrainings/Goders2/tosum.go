package main

import (
	"fmt"
)

func main() {
	// Method 1: Using a for loop
	fmt.Println("Method 1: Using a for loop")
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Method 2: Using a while loop
	fmt.Println("\nMethod 2: Using a while loop")
	sum = 0
	i := 1
	for i <= 10 {
		sum += i
		i++
	}
	fmt.Println(sum)

	// Method 3: Using a range loop
	fmt.Println("\nMethod 3: Using a range loop")
	sum = 0
	for i := range [10]int{} {
		sum += i + 1
	}
	fmt.Println(sum)

	// Method 4: Using the sum function from the math package
	fmt.Println("\nMethod 4: Using a loop")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var top int
	for _, num := range numbers {
		top += num
	}
	fmt.Println(top)

	// ex 5:
	fmt.Println("\nMethod 5: with a map")
	grades := map[string]int{"Alice": 90, "Bob": 80, "Charlie": 70}
	total := 0
	for _, grade := range grades {
		total += grade
	}
	fmt.Println(total)
}
