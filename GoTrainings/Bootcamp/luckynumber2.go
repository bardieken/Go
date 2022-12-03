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
`
	sec   = "You are good"
	tird  = "You are great"
	forth = "You are awesome"
	fifth = "You are a genius"
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
	if guess <= 0 {
		fmt.Println("Please provide a positive number")
		return
	}
	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(guess) + 1
		if n == guess {
			if turn == 1 {
				fmt.Println(" 1. time guessed right!")
			} else {
				switch turn {
				case 2:
					fmt.Println(sec)
				case 3:
					fmt.Println(tird)
				case 4:
					fmt.Println(forth)
				case 5:
					fmt.Println(fifth)
				}
			}
			return
		}
	}
	print := "%s Not Lucky %d times tried"
	switch rand.Intn(2) {
	case 0:
		fmt.Printf(print, "You lose!", maxTurns)
	case 1:
		fmt.Printf(print, "Try again!", maxTurns)
	}
}
