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
	str := os.Args[1]
	runes := []rune(str)
	for j := range runes {
		if runes[j] >= 'A' && runes[j] <= 'Z' {
			z01.PrintRune(155 - runes[j])
		} else if runes[j] >= 'a' && runes[j] <= 'z' {
			z01.PrintRune(219 - runes[j])
		} else {
			z01.PrintRune(runes[j])
		}
	}
	z01.PrintRune('\n')
}
