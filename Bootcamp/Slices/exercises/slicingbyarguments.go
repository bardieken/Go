package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	fmt.Printf("%q\n\n", ships)
	from, to := 0, len(ships) // pek anlamadim
	poss := os.Args[1:]
	if len(poss) == 0 {
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	} else if len(poss) == 1 {
		from, _ = strconv.Atoi(poss[0])
	} else if len(poss) == 2 {
		from, _ = strconv.Atoi(poss[0])
		to, _ = strconv.Atoi(poss[1])
	}
	if l := len(ships); from < 0 || from > l || to > l || from > to {
		fmt.Println("Wrong positions")
		return
	}

	fmt.Println(ships[from:to])
}
