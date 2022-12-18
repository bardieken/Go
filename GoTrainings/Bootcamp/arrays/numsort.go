package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	nums := []int{} // creates a slice of integers and initializes it with an empty slice
	var count int   // creates a variable with the type []int, which is a slice of integers
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%s Not a number skipping", arg)
			continue
		}
		nums = append(nums, num) // add num to nums slice
		count++
		if count > 5 {
			fmt.Printf("Max 5 numbers. Given %d numbers", count)
			break
		}
	}
	sort.Ints(nums) // sort the slice
	fmt.Println(nums)
}
