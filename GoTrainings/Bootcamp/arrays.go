package main

import "fmt"

func main() {
	/*myArray := [3]int{}
	myArray[0] = 32
	myArray[1] = 23
	myArray[2] = 54
	fmt.Println(myArray)
	*/
	/*var numbers = [5]int{4, 5, 7, 8, 9}
	fmt.Println(numbers)
	fmt.Println("There are", len(numbers), " numbers in the array")
	*/
	var cars [3]string
	cars[0], cars[1], cars[2] = "tesla", "mercedes", "jaguar"
	i := 0
	for i < len(cars) {
		fmt.Println(cars[i])
		i++
	}
}
