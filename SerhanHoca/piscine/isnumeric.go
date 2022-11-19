package piscine

func IsNumeric(s string) bool {
	word := []rune(s)
	counter := 0
	for i := range word {
		if word[i] >= '0' && word[i] <= '9' {
			counter++
		}
	}
	if counter == len(s) {
		return true
	}
	return false
}
