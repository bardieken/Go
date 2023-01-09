package main

import "fmt"

func main() {
	alphabet := []string{}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}
	fmt.Println(alphabet)
	my_slice_1 := []string{"Geeks", "for", "Geeks"}

	fmt.Println("My Slice 1:", my_slice_1)

	// Creating a slice
	// using shorthand declaration
	my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("My Slice 2:", my_slice_2)
	nums := make([]int, 10)
	for i := range nums {
		// do something with index
		fmt.Printf("%v \n", i)
	}
	var a [11]int
	for i := 1; i < len(a); i++ {
		a[i] = i
	}

	fmt.Print(a)
	fmt.Println(a[1:5])
}
