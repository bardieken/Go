package main

import "fmt"

func main() {
	sum := 0
	for i := 1; ; i++ {
		// i değeri 2 ve 3'e tam bölünüyorsa atla
		if i%2 == 0 && i%3 == 0 {
			continue
		}

		// sum değeri 15'ten büyükse döngüyü kır
		if sum > 15 {
			break
		}
		sum += i
	}
	fmt.Println(sum)
}
