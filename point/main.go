package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintNbr(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintNbr(points.y)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	ans := ""
	if n >= 0 {
		if n == 0 {
			ans += "0"
		}
		for n > 0 {
			last := n % 10
			ans = string('0'+rune(last)) + ans
			n = n / 10
		}
		for _, c := range ans {
			z01.PrintRune(c)
		}
	} else {
		for n <= -1 {
			last := n % 10
			last *= -1
			ans = string('0'+rune(last)) + ans
			n = n / 10
		}
		ans = "-" + ans
		for _, c := range ans {
			z01.PrintRune(c)
		}
	}
}
