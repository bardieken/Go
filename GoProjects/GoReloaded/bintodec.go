package main

import (
	"fmt"
	"strconv"
)

func main() {
	bin_num := "101010"
	dec_num, err := strconv.ParseInt(bin_num, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("decimal number is:", dec_num)
}

/* This function converts binary to decimal
func binToDec(binString string) string {
	output, err := strconv.ParseInt(binString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(int(output))
*/
