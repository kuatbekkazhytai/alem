package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	str := ""
	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	} else {
		str1 := args[0]
		str2 := args[1]
		a := []rune(str1)
		b := []rune(str2)

		for _, j := range a {
			for _, l := range b {
				if j == l {
					str = str + string(j)
				}
			}
		}
		fmt.Println(str)
		s := ""
		for _, j := range str {
			if isUnique(j, []rune(s)) {
				s = s + string(j)
			}
		}
		fmt.Println(s)
	}
}

func isUnique(r rune, a []rune) bool {
	for _, j := range a {
		if r == j {
			return false
		}
	}
	return true
}
