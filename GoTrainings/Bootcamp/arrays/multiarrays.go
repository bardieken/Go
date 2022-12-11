package main

import "fmt"

func main() {
	students := [][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}
	var sum float64
	for _, student := range students {
		for _, grade := range student {
			sum += grade
		}
		fmt.Println(sum)
	}
	N := float64(len(students) * len(students[0]))
	fmt.Println(N)
	fmt.Printf("Average: %0.2f\n", sum/N)
}
