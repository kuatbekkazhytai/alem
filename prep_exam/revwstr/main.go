package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
	} else {
		str := args[0]
		fmt.Println(str)
		var res string
		for i := len(str) - 1; i >= 0; i-- {
			if str[i] == ' ' {
				res = res + str[i+1:] + " "
				str = str[:i]
			}

		}
		fmt.Println(res)
		fmt.Println(str)

		res = res + str
		fmt.Println(res)
	}
}
