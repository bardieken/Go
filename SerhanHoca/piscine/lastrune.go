package piscine

func LastRune(s string) rune {
	lastLetter := []rune(s)
	lenght := len([]rune(s))
	return lastLetter[lenght-1]
}
