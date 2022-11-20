package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func nrToString(nr int) string {
	numberString := ""
	divider := 10
	digit := 0
	for nr > 0 {
		digit = nr % divider
		nr = nr / divider
		numberString = string(digit+48) + numberString
	}
	return numberString
}

func main() {
	points := &point{}
	setPoint(points)
	x := nrToString(points.x)
	y := nrToString(points.y)
	text := ""
	text = text + "x = " + x + ", y = " + y
	for _, element := range text {
		z01.PrintRune(element)
	}
	z01.PrintRune('\n')
}
