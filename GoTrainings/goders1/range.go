package main

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := 0
	for _, v := range n {
		t += v
	}
	println(t)
}

/*
func toplam(n []int]) int {
    t := 0

    for _, v := range n {
        t += v
    }

    return t
}

func toplam(n ...int) int {
    t := 0

    for _, v := range n {
      t += v
    }

    return t
}
*/
