package piscine

import "github.com/01-edu/z01"

var comb []int

func pnt(v []int) {
	for i := 0; i < len(v); i++ {
		z01.PrintRune('0' + rune(v[i]))
	}
	for j := 10 - len(v); j < 10; j++ {
		if j != v[j-(10-len(v))] {
			z01.PrintRune(',')
			z01.PrintRune(' ')
			return
		}
	}
}

func rec(cnt, n int) {
	if n == 0 {
		pnt(comb)
		return
	}
	for i := cnt; i <= 10-n; i++ {
		comb = append(comb, i)
		rec(i+1, n-1)
		comb = comb[:len(comb)-1]
	}
}

func PrintCombN(n int) {
	rec(0, n)
	z01.PrintRune('\n')
}
