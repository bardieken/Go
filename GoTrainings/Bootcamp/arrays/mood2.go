package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	moods := [][]string{
		{"awsome", "happy", "love", "great", "thankful"},
		{"sad", "terrible", "awful", "hopless", "depressed"},
	}
	// Print a random mood
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Please write your name and mood")
		return
	}
	name, mod := args[0], args[1]
	rand.Seed(time.Now().UnixNano())
	switch mod {
	case "positive":
		moodtext := moods[0][rand.Intn(len(moods[0]))]
		fmt.Printf("%s feels %s\n", name, moodtext)
	case "negative":
		moodtext := moods[1][rand.Intn(len(moods[1]))]
		fmt.Printf("%s feels %s\n", name, moodtext)

	}
}
