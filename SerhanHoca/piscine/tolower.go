package piscine

func ToLower(s string) string {
	word := []rune(s)
	counter := 0
	result := ""
	for i := range word {
		if word[i] >= 'A' && word[i] <= 'Z' {
			counter++
			word[i] = word[i] + 32
		}
		result += string(word[i])
	}
	return result
}
