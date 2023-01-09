package main

import "fmt"

func main() {
	x := 'a'
	for {
		if x > 'y' {
			break
		}
		fmt.Printf("%v,", string(x)) // demek ki rune ve string arasi print'de geciyor

		x++
	}
	fmt.Printf("z\n")
}
