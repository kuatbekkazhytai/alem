package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func main() {
	count := 0
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	a, _ := strconv.Atoi(os.Args[1])

	for i := 0; i <= a; i++ {
		if IsPrime(i) {
			count = count + i
		}
	}
	fmt.Println(count)
}

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	//var i
	for i := 3; i*i <= nb; i = i + 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
