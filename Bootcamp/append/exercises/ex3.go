package main

import "fmt"

func main() {
	toppings := []string{"black olives", "green peppers"}

	var pizza []string
	pizza = append(pizza, toppings...)
	pizza = append(toppings, "onions")
	toppings = append(pizza, "extra cheese")

	fmt.Printf("pizza       : %s\n", toppings)
}
