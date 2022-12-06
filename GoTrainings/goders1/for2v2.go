package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 4, 5, 4, 3, 2, 1, 0}
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}
}
