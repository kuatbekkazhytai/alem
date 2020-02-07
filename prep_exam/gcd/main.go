package main

import (
	
	"fmt"
	"os"
	"strconv"
)
func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println('\n')
	} else {
		num1, _ := strconv.Atoi(args[0])
		num2, _ := strconv.Atoi(args[1])
		for i := num1; i >= 1; i -- {
			if num1% i == 0 && num2%i == 0 {
				fmt.Println(i)
				return
			} 
		}
	} 
} 