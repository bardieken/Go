package main

import (
	"fmt"
	"strings"
)

func main() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay
	// yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday
	lyric = append([]string{"yesterday"}, lyric...)

	const N, M = 8, 13
	lyric = append(lyric, lyric[N:M]...)
	lyric = append(lyric[:N], lyric[M:]...)

	fmt.Printf("%s\n", lyric)
}
