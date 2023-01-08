package main

import "fmt"

func main() {
	data := []float64{10, 25, 30, 50}
	newData := []float64{99, 100}

	for i := range newData {
		data[i] = newData[i]
	}
	fmt.Println(data)
}
