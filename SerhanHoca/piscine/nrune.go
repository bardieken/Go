package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	len := 0
	for index := range a {
		len = index
	}
	if n-1 >= 0 && n-1 <= len {
		return a[n-1]
	}
	return 0
}
