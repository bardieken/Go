package main

import "fmt"

func main() {
	week := [4]string{"Monday", "Tuesday"}
	wend := [4]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	evens := [...]int{2, 4, 6, 8, 10}
	evens2 := [...]int32{2, 4, 6, 8, 10}
	evens2conv := [5]int{}
	for i := range evens2 {
		evens2conv[i] = int(evens2[i])
	}

	fmt.Println(evens == evens2conv)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.
	image := [5]uint8{'h', 'i'}
	var data [5]byte
	data = image

	fmt.Println(data == image)
}
