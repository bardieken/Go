package main

import "fmt"

func main() {

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	for i := 0; i < 3; i++ { // start of the execution block
		fmt.Println("Hello")
	}

	var items []int = []int{1, 2, 3, 4, 5}
	for i, v := range items {
		fmt.Println(i, v)

	}

	var item []int = []int{1, 2, 3, 4, 5}
	for _, v := range item {
		fmt.Println(v)
	}

	var m map[int]string = map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, j)
		}
	}
outside: // declare the label
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i < j {
				break outside // break to that label
			}
			fmt.Println(i, j)
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i < j {
				continue
			}
			fmt.Println(i, j)
		}
	}

}
