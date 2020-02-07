package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}
	str := args[0]
	runes := []rune(str)

	for _, j := range runes {
		if j >= 'a' && j <= 'm' {
			z01.PrintRune(j + 13)
		} else if j >= 'n' && j <= 'z' {
			z01.PrintRune(j - 13)
		} else if j >= 'A' && j <= 'M' {
			z01.PrintRune(j + 13)
		} else if j >= 'N' && j <= 'Z' {
			z01.PrintRune(j - 13)
		} else {
			z01.PrintRune(j)
		}
	}
	z01.PrintRune('\n')

}
