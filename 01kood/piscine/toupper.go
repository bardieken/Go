package piscine

func ToUpper(s string) string {
	word := []rune(s)
	counter := 0
	result := ""
	for i := range word {
		if word[i] >= 'a' && word[i] <= 'z' {
			counter++
			word[i] = word[i] - 32
		}
		result += string(word[i])
	}
	return result
}
