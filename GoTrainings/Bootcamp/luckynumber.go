package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 3
	usage    = `You have %d turns to guess the lucky number.`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Please provide a number")
		return
	}

	if guess < 0 {
		fmt.Println("Please provide a positive number")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(10) + 1
		if n == guess {
			fmt.Println("You win!")
			return
		}
	}
	fmt.Println("You lose!")
}
