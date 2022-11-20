package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		content, err := ioutil.ReadFile(os.Stdin.Name())
		if err != nil {
			PrintStrln("ERROR: open " + args[0] + ": no such file or directory")
			os.Exit(1)
		} else {
			PrintStr(string(content))
		}
	} else {
		for i := 0; i < len(args); i++ {
			content, err := ioutil.ReadFile(args[i])
			if err != nil {
				PrintStrln("ERROR: open " + args[i] + ": no such file or directory")
				os.Exit(1)
			} else {
				PrintStr(string(content))
			}
		}
	}
}

func PrintStrln(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func PrintStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
