package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		z01.PrintRune('\n')
	} else {
		runes1 := []rune(args[0])
		runes2 := []rune(args[1])
		str1 := ""
		str2 := ""
		for _, j := range runes1 {
			if IsUnique(j, []rune(str1)) {
				str1 = str1 + string(j)
			}
		}
		for _, j := range runes2 {
			if IsUnique(j, []rune(str2)) {
				str2 = str2 + string(j)
			}
		}
		str := str1 + str2
		res := ""
		for _, j := range []rune(str) {
			if IsUnique(j, []rune(res)) {
				res = res + string(j)
			}
		}
		fmt.Println(res)
	}
}

func IsUnique(r rune, runes []rune) bool {
	for _, j := range runes {
		if r == j {
			return false
		}
	}
	return true
}