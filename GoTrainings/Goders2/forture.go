package main

import "fmt"

func main() {
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++ // x++ printin ustune yazinca bir ekleyip basliyor!
	}
	fmt.Println("Done")
}
