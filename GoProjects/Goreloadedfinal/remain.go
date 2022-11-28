package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	// Reads the arguments from the command line
	givenText, _ := os.ReadFile(args[0])
	// Reads the first argument as a file
	// words := strings.Split(strings.Replace(string(givenText), "\n", "", -1), " ")
	words := strings.Split(string(givenText), " ")
	// Splits the text into words by dividing at spaces and stores it in a slice
	for i, word := range words {
		switch word {
		case "(up)": // If the word is (up), replace the word before it with the uppercase version of the word
			words[i-1] = strings.ToUpper(words[i-1])  // Uppercases the word before (up)
			words = append(words[:i], words[i+1:]...) // Removes (up) from the slice by appending the slice before (up) and after (up)
		case "(down)": // If the word is (down), replace the word before it with the lowercase version of the word
			words[i-1] = strings.ToLower(words[i-1])  // Lowercases the word before (down)
			words = append(words[:i], words[i+1:]...) // Removes (down) from the slice by appending the slice before (down) and after (down)
		case "(cap)": // If the word is (cap), replace the word before it with the capitalized version of the word
			words[i-1] = strings.Title(words[i-1])    // Capitalizes the word before (cap)
			words = append(words[:i], words[i+1:]...) // Removes (cap) from the slice by appending the slice before (cap) and after (cap)
		case "(hex)": // If the word is (hex), replace the word before it with the integer version of the hexadecimal number
			words[i-1] = HextoInt(words[i-1])         // Converts the hexadecimal number before (hex) to an integer
			words = append(words[:i], words[i+1:]...) // Removes (hex) from the slice by appending the slice before (hex) and after (hex)
		case "(bin)": // If the word is (bin), replace the word before it with the integer version of the binary number
			words[i-1] = BintoInt(words[i-1])         // Converts the binary number before (bin) to an integer
			words = append(words[:i], words[i+1:]...) // Removes (bin) from the slice by appending the slice before (bin) and after (bin)
		case "(up,": // If the word is (up, n), replace the word before it with the uppercase version of the word, but only the first n letters
			b := strings.Trim(string(words[i+1]), words[i+1][1:]) // Removes the last character of the word after (up,
			number, _ := strconv.Atoi(string(b))                  // Converts the number after (up, to an integer
			for j := 1; j <= number; j++ {                        // Loops through the word before (up, n) and capitalizes the first n letters
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...) // Removes (up, n) from the slice by appending the slice before (up, n) and after (up, n)
		case "(low,": // If the word is (low, n), replace the word before it with the lowercase version of the word, but only the first n letters
			b := strings.Trim(string(words[i+1]), words[i+1][1:]) // Removes the last character of the word after (low,
			number, _ := strconv.Atoi(string(b))                  // Converts the number after (low, to an integer
			for j := 1; j <= number; j++ {                        // Loops through the word before (low, n) and lowercases the first n letters
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...) // Removes (low, n) from the slice by appending the slice before (low, n) and after (low, n)
		case "(down,": // If the word is (down, n), replace the word before it with the lowercase version of the word, but only the first n letters
			b := strings.Trim(string(words[i+1]), words[i+1][1:]) // Removes the last character of the word after (down,
			number, _ := strconv.Atoi(string(b))                  // Converts the number after (down, to an integer
			for j := 1; j <= number; j++ {                        // Loops through the word before (down, n) and lowercases the first n letters
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...) // Removes (down, n) from the slice by appending the slice before (down, n) and after (down, n)
		case "(cap,": // If the word is (cap, n), replace the word before it with the capitalized version of the word, but only the first n letters
			b := strings.Trim(string(words[i+1]), words[i+1][1:]) // Removes the last character of the word after (cap,
			number, _ := strconv.Atoi(string(b))                  // Converts the number after (cap, to an integer
			for j := 1; j <= number; j++ {                        // Loops through the word before (cap, n) and capitalizes the first n letters
				words[i-j] = strings.Title(words[i-j])
			}
			words = append(words[:i], words[i+2:]...) // Removes (cap, n) from the slice by appending the slice before (cap, n) and after (cap, n)
		}
	}
	// Change "a" to "an" if the word starts with a vowel
	ChangeA(words)
	// Joins the slice into a string and adds punctuation
	needed := strings.Join(Punctuations(words), " ")
	// Writes the string to the file specified by the second argument with 644 permissions
	man := os.WriteFile(args[1], []byte(needed), 0o644)
	if man != nil {
		panic(man) // If there is an error, panic
	}
}

// Converts hexadecimal numbers to integers
func HextoInt(hex string) string {
	number, _ := strconv.ParseInt(hex, 16, 64) // Converts the hexadecimal number to an integer by parsing it as a base 16 number
	return fmt.Sprint(number)                  // Returns the integer as a string
}

// Converts binary numbers to integers
func BintoInt(bin string) string {
	number, _ := strconv.ParseInt(bin, 2, 64) // Converts the binary number to an integer by parsing it as a base 2 number
	return fmt.Sprint(number)                 // Returns the integer as a string
}

// Change "a" to "an" if the word starts with a vowel
func ChangeA(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"} // Creates a slice of vowels
	for i, word := range s {                                                       // Loops through the slice
		for _, letter := range vowels { // Loops through the slice of vowels
			// if the word starts with a vowel, change "a" to "an"
			if word == "a" && string(s[i+1][0]) == letter { // If the word is "a" and the next word starts with a vowel
				s[i] = "an"
			} else if word == "A" && string(s[i+1][0]) == letter { // If the word is "A" and the next word starts with a vowel
				s[i] = "An"
			}
		}
	}
	return s // Returns the new slice
}

// Adds punctuation
func Punctuations(s []string) []string {
	puncs := []string{",", ".", "!", "?", ":", ";"} // Creates a slice of punctuation
	// punc in the middle of a string connecting to word after
	for i, word := range s { // Loops through the slice
		for _, punc := range puncs { // Loops through the slice of punctuation
			if string(word[0]) == punc && string(word[len(word)-1]) != punc { // if the word starts with a punctuation, add the punctuation to the end of the word before it
				s[i-1] += punc  // Adds the punctuation to the end of the word before it
				s[i] = word[1:] // Removes the punctuation from the beginning of the word
			}
		}
	}
	// punc at end of string
	for i, word := range s { // Loops through the slice
		for _, punc := range puncs { // Loops through the slice of punctuation
			if (string(word[0]) == punc) && (s[len(s)-1] == s[i]) { // if the word starts with a punctuation, add the punctuation to the end of the last word
				s[i-1] += word   // Adds the punctuation to the end of the last word
				s = s[:len(s)-1] // Removes the punctuation from the slice
			}
		}
	}
	// punc in middle of string
	for i, word := range s { // Loops through the slice
		for _, punc := range puncs { // Loops through the slice of punctuation
			if string(word[0]) == punc && string(word[len(word)-1]) == punc && s[i] != s[len(s)-1] { // if the word starts and ends with a punctuation, add the punctuation to the end of the word before it
				s[i-1] += word                // Adds the punctuation to the end of the word before it
				s = append(s[:i], s[i+1:]...) // Removes the punctuation from the slice
			}
		}
	}
	// for apostrophe
	count := 0               // Creates a counter
	for i, word := range s { // Loops through the slice
		if word == "'" && count == 0 { // if the word is an apostrophe and the counter is 0, add the apostrophe to the end of the word before it
			count += 1                    // Adds 1 to the counter
			s[i+1] = word + s[i+1]        // Adds the apostrophe to the end of the word before it
			s = append(s[:i], s[i+1:]...) // Removes the apostrophe from the slice
		}
	}
	//  for second apostrophe
	for i, word := range s { // Loops through the slice
		if word == "'" { // if the word is an apostrophe, add the apostrophe to the end of the word before it
			// print("here")
			s[i-1] = s[i-1] + word // Adds the apostrophe to the end of the word before it
			// print(s[i-1])
			s = append(s[:i], s[i+1:]...) // Removes the apostrophe from the slice
		}
	}
	return s // Returns the new slice
}
