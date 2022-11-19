package piscine

func IsPrintable(s string) bool {
	word := []rune(s)
	counter := 0
	for i := range word {
		if word[i] > 31 && word[i] < 128 {
			counter++
		}
	}
	if counter == len(s) {
		return true
	}
	return false
}
