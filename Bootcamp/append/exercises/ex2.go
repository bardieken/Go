package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		PizzaTtoppings []string
		DepartureTimes []time.Time
		Graduations    []int
		onOff          []bool
	)
	now := time.Now()

	PizzaTtoppings = append(PizzaTtoppings, "ham", "cheese", "salam")
	hours := []int{3, 6, 9}
	for i := range hours {
		// for i := 0; i < 8; i++ {
		DepartureTimes = append(DepartureTimes, now.Add(time.Hour*time.Duration(hours[i])))
	}
	Graduations = append(Graduations, 1998, 2005, 2018)
	onOff = append(onOff, true, false, true)
	fmt.Println(PizzaTtoppings)
	fmt.Println(DepartureTimes)
	fmt.Println(Graduations)
	fmt.Println(onOff)
}
