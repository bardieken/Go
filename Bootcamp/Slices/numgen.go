package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	max, _ := strconv.Atoi(os.Args[1])
	var uniq []int
loop:
	for len(uniq) < max {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")
		for _, u := range uniq {
			if u == n {
				continue loop
			}
		}
		uniq = append(uniq, n)

	}
	fmt.Println("\n\nuniques:", uniq)
	sort.Ints(uniq)
	fmt.Println("\nsorted", uniq)
}
