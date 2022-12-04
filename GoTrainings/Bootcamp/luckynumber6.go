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
	min := 10
	if guess > min {
		min = guess
	}
	var turns int
	turns = guess / 4
	if turns < maxTurns {
		turns = maxTurns
	}
	for turn := 0; turn < turns; turn++ {
		n := rand.Intn(min) + 1
		println(n)
		if n == guess {

			fmt.Println("You guessed right!")
			return
		}
	}
	fmt.Println("You lose!")
}
