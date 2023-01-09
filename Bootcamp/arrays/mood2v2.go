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
	n := rand.Intn(len(moods[0]))

	var mi int
	if mod != "positive" {
		mi = 1
	}
	fmt.Printf("%s feels %s\n", name, moods[mi][n])
}

//explanation:
/*moods := [][]string{
    {"awsome", "happy", "love", "great", "thankful"},
    {"sad", "terrible", "awful", "hopless", "depressed"},
}

name := "Bob"
mod := "negative"

var mi int
if mod != "positive" {
    mi = 1
}

// Generate a random number between 0 and the length of the
// selected moods array (which will be 1 in this case)
rand.Seed(time.Now().UnixNano())
n := rand.Intn(len(moods[mi]))

// Print a random mood from the selected moods array
fmt.Printf("%s feels %s\n", name, moods[mi][n])
*/
