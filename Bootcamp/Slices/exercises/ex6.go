package main

import (
	"fmt"
	"strings"
)

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
	for i := 0; i < len(namesB); i++ {
		k := strings.Split(namesA, ",")
		l := strings.Split(namesB[i], ",")

		if len(k) == len(l) {
			return
		}

	}
	fmt.Println("They are Equal")
}
