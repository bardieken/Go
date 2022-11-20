package piscine

func Fibonacci(index int) int {
	a := index
	if a < 0 {
		return -1
	}
	if a == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	return Fibonacci(a-1) + Fibonacci(a-2)
}
