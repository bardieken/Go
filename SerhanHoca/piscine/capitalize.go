package piscine

// Write a function that capitalizes the first letter of each word and lowercases the rest.

// A word is a sequence of alphanumeric characters.

func Capitalize(s string) string {
	runes := []rune(s)

	first := true

	for i := range runes {
		if isNumberOrLetter(runes[i]) == true && first {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = runes[i] - 32
			}
			first = false
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] = runes[i] + 32
		} else if isNumberOrLetter(runes[i]) == false {
			first = true
		}
	}
	return string(runes)
}

func isNumberOrLetter(char rune) bool {
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
		return true
	}
	return false
}
