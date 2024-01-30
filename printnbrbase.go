package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	res := ""
	sign := ""

	if checkbase(base) {
		d := len(base)
		if nbr < 0 {
			sign = "-"
			for nbr < 0 {
				rem := nbr % d
				nbr /= d
				res = string(base[-rem]) + res
			}
		} else {
			for nbr > 0 {
				rem := nbr % d
				nbr /= d
				res = string(base[rem]) + res
			}
		}

	} else {
		res += "NV"
	}
	res = sign + res
	for _, c := range res {
		z01.PrintRune(c)
	}
}

func checkbase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i, c := range base {
		if c == '+' || c == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if rune(base[j]) == c {
				return false
			}
		}
	}
	return true
}
