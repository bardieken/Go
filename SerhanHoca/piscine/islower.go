package piscine

func IsLower(s string) bool {
	word := []rune(s)
	counter := 0
	for i := range word {
		if word[i] >= 'a' && word[i] <= 'z' {
			counter++
		}
	}
	if counter == len(s) {
		return true
	}
	return false
}
