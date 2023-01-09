package main

import "fmt"

func main() {
	names := [3]string{"Ali", "Veli", "Selami"}
	distances := [5]int{10, 20, 30, 40, 50}
	data := [5]byte{'h', 'e', 'l', 'L', 'o'}
	ratios := [1]float64{3.14}
	alives := [4]bool{true, false, true, true}
	zero := []byte{}

	for i := 0; i < len(names); i++ {
		fmt.Printf(" Best friends are : %q\n", names[i])
	}
	for i := 0; i < len(distances); i++ {
		fmt.Printf(" Distances are : %d\n", distances[i])
	}
	for i := 0; i < len(data); i++ {
		fmt.Printf(" Data are : %d\n", data[i])
	}
	for i := 0; i < len(ratios); i++ {
		fmt.Printf(" Ratios are : %f\n", ratios[i])
	}
	for i := 0; i < len(alives); i++ {
		fmt.Printf(" Alives are : %t\n", alives[i])
	}
	for i := 0; i < len(zero); i++ {
		fmt.Printf(" Zero are : %d\n", zero[i])
	}
	for i := range names {
		fmt.Printf(" Best Friends are : %q\n", names[i])
	}
	for i := range distances {
		fmt.Printf(" Distances are : %d\n", distances[i])
	}
	for i := range data {
		fmt.Printf(" Data are : %d\n", data[i])
	}
	for i := range ratios {
		fmt.Printf(" Ratios are : %f\n", ratios[i])
	}
	for i := range alives {
		fmt.Printf(" Alives are : %#v\n", alives[i])
	}
	fmt.Print("\ndata")
	for i := 0; i < len(data); i++ {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}
}
