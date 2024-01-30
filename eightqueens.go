package piscine

import (
	"github.com/01-edu/z01"
)

func EightQueens() {
	x := []int{0, 0, 0, 0, 0, 0, 0, 0}
	k := 0
	chess(x, k)
}

func check(x []int, k int, j int) bool {
	for i := 0; i < k; i++ {
		if x[i] == j || i-x[i] == k-j || i+x[i] == k+j {
			return false
		}
	}
	return true
}

func chess(x []int, k int) {
	if k == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(x[i]+1) + '0')
		}
		z01.PrintRune('\n')
		return
	}
	if k < 0 {
		return
	}
	for i := 0; i < 8; i++ {
		if check(x, k, i) {
			x[k] = i
			chess(x, k+1)
		}
	}
}
