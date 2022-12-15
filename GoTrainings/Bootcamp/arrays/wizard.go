package main

import "fmt"

func main() {
	wizards := [][]string{
		{"Albert", "Enstain", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "black hole"},
		{"Marie", "Curie", "radioactivity"},
		{"Charles", "Darwin", "evolution"},
	}

	n, l, k := "First Name", "Last Name", "Nickname"

	fmt.Printf("%-15s %-15s %-15s\n", n, l, k)
	for i := 0; i < 50; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
	// for _, wizard := range wizards {
	// fmt.Printf("%-15s %-15s %-15s\n", wizard[0], wizard[1], wizard[2])
	for i := 0; i < len(wizards); i++ {
		fmt.Printf("%-15s %-15s %-15s\n", wizards[i][0], wizards[i][1], wizards[i][2])
	}
}
