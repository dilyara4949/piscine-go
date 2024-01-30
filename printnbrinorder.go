package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	ans := ""

	if n == 0 {
		ans += "0"
	}
	for n > 0 {
		last := n % 10
		ans = string('0'+rune(last)) + ans
		n = n / 10
	}
	m := len(ans)
	anss := []rune(ans)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			for k := j + 1; k < m; k++ {
				if anss[j] > anss[k] {
					tmp := anss[j]
					anss[j] = anss[k]
					anss[k] = tmp
				}
			}
		}
	}
	for _, c := range anss {
		z01.PrintRune(c)
	}
}
