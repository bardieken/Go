package main

import "fmt"

func main() {

	/*myArray := []int{1, 2, 3, 4, 5, 6, 7}
	mySlice := myArray[:]
	fmt.Println(mySlice)
	mySlice[6] = 11
	fmt.Println(mySlice)
	*/
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	colors = append(colors[1 : len(colors)-1])
	fmt.Println(colors)
}
