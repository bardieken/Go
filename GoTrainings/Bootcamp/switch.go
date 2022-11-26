package main

import (
	"fmt"
	"os"
)

func main() {
	// city := "Tokyo"
	city := os.Args[1]

	switch city { // condition Expression string ise case clause de int olur
	case "Paris", "Lyon", "Marseille": // virgul ile birden fazla case (condition) yazılabilir
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where is this city?")
	}
}
