package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		z01.PrintRune('\n')
	}

	if !IsNumber(args[0]) || !IsValidSign(args[1]) || !IsNumber(args[2]) {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	num1, _ := strconv.Atoi(args[0])
	num2, _ := strconv.Atoi(args[2])
	if args[1] == "+" {
		fmt.Println(Addition(num1, num2))
	}
	if args[1] == "-" {
		fmt.Println(Subtraction(num1, num2))
	}
	if args[1] == "*" {
		fmt.Println(Multiplication(num1, num2))
	}
	if args[1] == "/" {

		if args[2] == "0" {
			PrintStr("No division by 0")
		} else {
			fmt.Println(Division(num1, num2))
		}

	}
	if args[1] == "%" {
		if args[2] == "0" {
			PrintStr("No modulo by 0")
		} else {
			fmt.Println(Modulo(num1, num2))
		}

	}

}

func IsNumber(s string) bool {
	for _, j := range s {
		if j >= '0' && j <= '9' {
			return true
		}
	}
	return false

}

func IsValidSign(s string) bool {
	l := 0
	for range s {
		l++
	}
	if l == 1 {
		if s[0] == '+' || s[0] == '-' || s[0] == '/' || s[0] == '*' || s[0] == '%' {
			return true
		}
	}
	return false
}

func Addition(a, b int) int {
	return a + b
}
func Subtraction(a, b int) int {
	return a - b
}
func Multiplication(a, b int) int {
	return a * b
}
func Division(a, b int) int {
	return a / b
}
func Modulo(a, b int) int {
	return a % b
}

func PrintStr(str string) {
	for _, j := range str {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}
