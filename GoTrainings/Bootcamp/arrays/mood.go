package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("Please write your name")
		return
	}
	name := args[0]
	mood := []string{"sad", "happy", "love", "hopless", "thankful"}
	rand.Seed(time.Now().UnixNano())
	moodtext := mood[rand.Intn(len(mood))]
	fmt.Printf("%s feels %s\n", name, moodtext)
	// Print a random mood
}
