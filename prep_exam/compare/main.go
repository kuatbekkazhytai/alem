package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		z01.PrintRune('\n')
	} else {
		if args[0] < args[1] {
			z01.PrintRune('-')
			z01.PrintRune('1')
			z01.PrintRune('\n')
		} else if args[0] > args[1] {
			z01.PrintRune('1')
			z01.PrintRune('\n')
		} else {
			z01.PrintRune('0')
			z01.PrintRune('\n')
		}
	}
}
