package main

import (
	"fmt"
	"os"
)

const (
	usage       = "Usage: <username> <password>"
	errUser     = "%q is not a username.\n "
	errPass     = "%q is wrong password. \n"
	accGranted  = "Access Granted: %q %q\n"
	user, user2 = "Ali", "Veli"
	pass, pass2 = "12345", "54321"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	u, p := args[1], args[2]
	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user && p == pass {
		fmt.Printf(accGranted, u, p)
	} else if u == user2 && p == pass2 {
		fmt.Printf(accGranted, u, p)
	} else {
		fmt.Printf(errPass, p)
	}
}
