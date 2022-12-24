package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if a := os.Args; len(a) != 2 {
		fmt.Printf("%q is not a valid month", a[1])
		return
	}
	// get the current year find leap
	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)
	days := 28
	month := strings.ToLower(os.Args[1])
	if month == "january" || month == "march" || month == "may" || month == "july" || month == "august" || month == "october" || month == "december" {
		days = 31
	} else if month == "april" || month == "june" || month == "september" || month == "november" {
		days = 30
	} else if month == "february" { // && leap {???? feb calismiyor
		if leap {
			days = 29
		}
	} else if month == "february" {
	} else {
		fmt.Printf("%q is not a valid month", month)
		return
	}
	fmt.Printf("%q has %d days.\n", month, days)
}
