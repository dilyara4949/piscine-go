package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	dif := 'a'
	i := 1
	if len(arg) > 1 && arg[1] == "--upper" {
		dif = 'A'
		i = 2
	}
	for ; i < len(arg); i++ {
		x := Atoi(arg[i])
		if x <= 0 || x > 26 {
			z01.PrintRune(' ')
		} else {
			z01.PrintRune(rune(x + int(dif) - 1))
		}
	}
	if len(arg) > 1 {
		z01.PrintRune('\n')
	}
}

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	sign := 0
	m := 1
	if s[0] == '-' {
		sign = -1
	} else if s[0] == '+' {
		sign = 1
	}
	end := 0
	if sign != 0 {
		end = 1
	}
	for i := len(s) - 1; i >= end; i-- {
		if rune(s[i])-'0' > 9 || rune(s[i])-'0' < 0 {
			return 0
		}
		res += m * int(rune(s[i])-'0')
		m *= 10
	}
	if sign == 0 {
		sign = 1
	}
	return res * sign
}
