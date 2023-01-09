package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
(Provide -v flag to see the picked numbers.)
`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}
	var verbose bool
	if args[0] == "-v" {
		verbose = true
	}
	guess, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Please provide a number")
		return
	}
	if guess <= 0 {
		fmt.Println("Please provide a positive number")
		return
	}
	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(guess) + 1
		if verbose {
			fmt.Printf("%d ", n)
		}
		if n == guess {
			if turn == 1 {
				fmt.Println(" 1. time guessed right!")
			} else {
				fmt.Printf("%d. time guessed right!\n", turn)
			}
			return
		}

	}
	fmt.Println("You lose!")
}
