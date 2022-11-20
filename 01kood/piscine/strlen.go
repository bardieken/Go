package piscine

func StrLen(s string) int {
	c := 0
	a := []rune(s)
	for range a {
		c++
	}
	return c
}
