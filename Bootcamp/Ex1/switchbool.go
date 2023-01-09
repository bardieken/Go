package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	number := os.Args[1]
	i, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}
	switch { // true yazmasak da olur default true
	case i > 0:
		fmt.Println("positive")
		// fallthrough yazilabilir
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")

	}
}

// fallthrough yazilinca bir sonraki case i control etmeden sonucu verir
// ilk case kontrol eder posotove veriri sonra fallthrough yazildigi icin negative yazar
