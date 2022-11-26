package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Hour() {
	case 5, 6, 7, 8, 9, 10, 11:
		fmt.Println("Good morning!")
	case 12, 13, 14, 15, 16:
		fmt.Println("Good afternoon.")
	case 17, 18, 19, 20:
		fmt.Println("Good evening.")
	case 21, 22, 23, 0, 1, 2, 3, 4:
		fmt.Println("Good night.")
	}
}
