package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}
	fmt.Println("Original:", items)
	// ADD YOUR CODE HERE
	sort.Slice(items[5:7], func(i, j int) bool { return items[5+i] < items[5+j] })
	fmt.Println()
	fmt.Println("Sorted  :", items)
}
