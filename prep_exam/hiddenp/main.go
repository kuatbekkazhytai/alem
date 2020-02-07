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
		str1 := args[0]
		r1 := []rune(str1)
		str2 := args[1]
		r2 := []rune(str2)
		count := 0

		for _, j := range r2 {

			if count == len(r1) {
				z01.PrintRune('1')
				z01.PrintRune('\n')
				return
			}
			if r1[count] == j {
				count++
			}
		}
		z01.PrintRune('0')
		z01.PrintRune('\n')

	}
}
