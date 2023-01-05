package main

import "fmt"

func main() {
	word := "numberofcharacters"
	a := make([]string, len(word))
	for i, r := range word {
		a[i] = string(r)
	}
	fmt.Println(a)
	fmt.Println(len(a))
}
