package main

import "fmt"

var (
	a = []string{
		"pacman1", "mario2", "tetris3", "doom4", "galaga5", "frogger6",
		"asteroids7", "simcity8", "metroid9", "defender10", "rayman11",
		"tempest12", "ultima13",
	} // 13 tane
	b []string
	c []string
	d []string
)

func main() {
	copy1()
	fmt.Println("copy1:", b)
	copy2()
	fmt.Println("copy2:", b)
	cut1()
	fmt.Println("Cut1:", d)
	del1()
	fmt.Println("del1:", c)
}

// Copy
func copy1() {
	b = make([]string, len(a))
	copy(b, a)
}

func copy2() {
	b = append(a[:0:0], a...) // a[:0:0] creates a new empty slice with the same length and capacity as a.
}

// CUT
func cut1() {
	d = make([]string, len(a))
	copy(d, a)
	d = append(d[:1], d[11:]...)
}

// Delete
func del1() {
	c = make([]string, len(a))
	copy(c, a)
	c = append(c[:2], c[3:]...)
}
