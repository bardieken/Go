package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var nums []int // cleate slice
	for _, arg := range os.Args[1:] {
		if num, err := strconv.Atoi(arg); err == nil { // checks conv. succesfull
			nums = append(nums, num) // add num to nums slice
		} else {
			nums = append(nums, 0) // append 0 instead non numeric values
		}
	}
	if len(nums) > 5 {
		// os.Exit(0)
		fmt.Printf("Max 5 numbers. Given %d numbers\n", len(nums))
		nums = nums[:5] // slice nums to only keep the first 5 elements
	}
	sort.Ints(nums) // sort the slice
	fmt.Println(nums)
}
