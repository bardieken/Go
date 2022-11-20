package piscine

func IsAlpha(s string) bool {
	word := []rune(s)
	counter := 0 // text
	for i := range word {
		if word[i] >= '0' && word[i] <= '9' || word[i] == ' ' || word[i] >= 'a' && word[i] <= 'z' || word[i] >= 'A' && word[i] <= 'Z' {
			counter++
		}
	}
	if counter == len(s) {
		return true
	}
	return false
}
