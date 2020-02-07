package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
	} else {
		runes := []rune(args[0])
		count := 0
		for i, j := range runes {
			if j >= 'a' && i <= 'z' {
				count = int(runes[i] - 96)
			} else if j >= 'A' && i <= 'Z' {
				count = int(runes[i] - 64)
			}
			for k := 0; k < count; k++ {
				z01.PrintRune(runes[i])
			}
		}

	}
}
