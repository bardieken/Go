package main

import (
	"fmt"
	"os"
)

const (
	usage      = "Usage: <username> <password>"
	errUser    = "%q is not a username.\n "
	errPass    = "%q is wrong password. \n"
	accGranted = "Access Granted: %q %q\n"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	u, p := args[1], args[2]
	if u != "Ali" {
		fmt.Printf(errUser, u)
	} else if p != "12345" {
		fmt.Printf(errPass, p)
	} else {
		fmt.Printf(accGranted, u, p)
	}
}
