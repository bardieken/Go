package main

import (
	"fmt"
	"strconv"
)

func main() {
	hex_num := "0x1a"
	dec_num, err := strconv.ParseInt(hex_num, 0, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("decimal number is:", dec_num)
}

/*func hexNumToInt(hexaString string) string {
	hexaString = strings.ReplaceAll(hexaString, "0x", "")
	hexaString = strings.ReplaceAll(hexaString, "0X,", "")
	output, err := strconv.ParseInt(hexaString, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(int(output))
}
*/
