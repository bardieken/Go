package main

import (
	"fmt"
	"strings"
)

func main() {
	tasks := []string{"jump", "run", "read"}
	upTasks := make([]string, 0, len(tasks)) // 0 length, len(tasks) capacity
	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
	}
	// pik := make([]string, len(tasks))
	pik := append([]string{"Today"}, tasks...)
	// copy(pik, []string "Today")
	fmt.Println(upTasks)
	fmt.Println(pik)
}
