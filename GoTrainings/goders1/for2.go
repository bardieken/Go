package main

func main() {
	var a []int = []int{1, 2, 4, 5, 4, 3, 2, 1, 0}

	for i, element := range a {
		// fmt.Println(i, element)
		for j, element2 := range a {
			if i < j && element == element2 {
				println(element)
			}
		}
	}
}
