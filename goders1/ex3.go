package main

import "fmt"

func main() {
	var (
		sayi [5]int
		tek  []int
		cift []int
		neg  []int
	)
	fmt.Println(" 5 sayi gir")
	for i := 0; i < 5; i++ {
		fmt.Scan(&sayi[i])
	}
	for i := 0; i < 5; i++ {
		if sayi[i]%2 == 0 && sayi[i] > 0 {
			cift = append(cift, sayi[i])
		} else if sayi[i]%2 != 0 && sayi[i] > 0 {
			tek = append(tek, sayi[i])
		} else {
			neg = append(neg, sayi[i])
		}
	}
	fmt.Println("Girilen sayilar", sayi)
	fmt.Println("Cift sayilar", cift)
	fmt.Println("Tek sayilar", tek)
	fmt.Println("negatif sayilar", neg)
}
