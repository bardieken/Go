package main

import (
	"fmt"
	"strings"
)

func main() {
	tasks := []string{"Jump", "run", "read"}
	var upTasks []string

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
	}
	fmt.Println(upTasks)
}
