package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int
	var count int
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%s is not a number, skipping...\n", arg)
			continue
		}
		sum += num
		count++
		if count == 5 {
			break
		}
	}
	fmt.Printf("Total: %d\n", sum)
	if count > 0 {
		fmt.Printf("Average: %.2f\n", float64(sum)/float64(count))
	} else {
		fmt.Println("No numbers were given.")
	}
}
