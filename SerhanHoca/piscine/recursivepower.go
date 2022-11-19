package piscine

func RecursivePower(nb int, power int) int {
	a := nb
	b := power
	if b < 0 {
		return 0
	} else {
		if b == 0 {
			return 1
		}
		return a * RecursivePower(a, b-1)
	}
}
