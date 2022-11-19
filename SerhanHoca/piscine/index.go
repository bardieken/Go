package piscine

func Index(s string, toFind string) int {
	str := []rune(s)
	find := []rune(toFind)
	n := len(str)
	k := len(find)
	for i := 0; i <= n-k; i++ {
		if toFind == s[i:i+k] {
			return (i)
		}
	}
	return -1
}
