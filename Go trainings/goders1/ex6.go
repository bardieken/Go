package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	prg_name := os.Args[0]
	fmt.Printf("The program name is %s\n", prg_name)

	names := os.Args[1:]
	fmt.Println(reflect.TypeOf(names))

	for _, name := range names {
		fmt.Printf("Hello, %s!\n", name)
	}
}
