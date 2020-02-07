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
		//runes := []rune(args[0])
		str := args[0]
		count := 0
		var arr []int
		for len(str) > 0 {
			count, str = Counter(rune(str[0]), str)
			arr = append(arr, count)
		}
		fmt.Println(arr)
		for i, j := range arr {
			for k, l := range arr {
				if j == l && i != k {
					fmt.Println(false)
					return
				}
			}
		}
		fmt.Println(true)
	}
}

func Counter(r rune, s string) (int, string) {
	count := 0
	res := ""
	for _, j := range s {
		if r == j {
			count++
		} else {
			res = res + string(j)
		}
	}
	return count, res
}
