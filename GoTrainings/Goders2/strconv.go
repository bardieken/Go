package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "2 4 6 1 3 5 a w 11 S s"
	fields := strings.Split(data, " ")
	var ints []int
	for _, field := range fields {
		i, err := strconv.Atoi(field)
		if err == nil {
			fmt.Println(i)
			ints = append(ints, i)
		}
	}
	fmt.Println(ints)
}
