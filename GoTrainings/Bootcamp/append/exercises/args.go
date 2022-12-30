package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var input []int
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Gief Numbers")
		return
	}
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			continue
		}
		input = append(input, num)
	}
	sort.Ints(input)
	fmt.Println(input)
}
