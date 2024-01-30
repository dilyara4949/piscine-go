package piscine

import "github.com/01-edu/z01"

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
