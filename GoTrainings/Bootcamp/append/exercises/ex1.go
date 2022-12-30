package main

import (
	"bytes"
	"fmt"
)

func main() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	// append elements to header to make it equal with the png slice
	// compare the slices using the bytes.Equal function
	// print whether they're equal or not

	header = append(header, png...)
	equal := bytes.Equal(png, header)
	if equal {
		fmt.Println("They are equal")
	} else {
		fmt.Println("Not equal")
	}
	// if bytes.Equal(png, header) {
	// fmt.Println("They are equal")
}
