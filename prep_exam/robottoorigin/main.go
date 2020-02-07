package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]

	r := 0
	l := 0
	u := 0
	d := 0
	if len(args) != 1 {
		z01.PrintRune('\n')
	} else {
		str := args[0]
		for _, j := range str {
			if j == 'R' {
				r++
			} else if j == 'L' {
				l++
			} else if j == 'U' {
				u++
			} else if j == 'D' {
				d++
			}
		}
		if r == l && r != 0 {
			z01.PrintRune('t')
			z01.PrintRune('r')
			z01.PrintRune('u')
			z01.PrintRune('e')
			z01.PrintRune('\n')
		} else if u == d && u != 0 {
			z01.PrintRune('t')
			z01.PrintRune('r')
			z01.PrintRune('u')
			z01.PrintRune('e')
			z01.PrintRune('\n')
		} else {
			z01.PrintRune('f')
			z01.PrintRune('a')
			z01.PrintRune('l')
			z01.PrintRune('s')
			z01.PrintRune('e')
			z01.PrintRune('\n')
		}
	}
}
