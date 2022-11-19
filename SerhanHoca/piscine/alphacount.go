package piscine

func AlphaCount(s string) int {
	sentence := []rune(s)
	counter := 0
	for i := range sentence {
		if sentence[i] >= 'A' && sentence[i] <= 'Z' || sentence[i] >= 'a' && sentence[i] <= 'z' {
			counter++
		}
	}
	return counter
}
