package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	x := 32
	for x < 122 {
		x++
		//if x == 39 {
		//continue
		//}
		h := fmt.Sprintf("%x\n", x)
		print, err := hex.DecodeString(h)
		if err != nil {
			fmt.Printf("index: %d \t %v\n", x, string(print))
		}
	}
}
