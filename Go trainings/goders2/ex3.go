package main

import "fmt"

func main() {
	a := [...]string{"abgun bakgun", "papgun hak", "burt jump"}
	for i := range a {
		fmt.Println("Array item:", i, "is", a[i])
	}
}
