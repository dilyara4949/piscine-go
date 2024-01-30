package piscine

func AtoiBase(s string, base string) int {
	res := 0
	if checkbase(base) {
		n := len(base)
		m := 1
		for i := len(s) - 1; i >= 0; i-- {
			idx := find(base, rune(s[i]))
			if idx == -1 {
				return 0
			}
			res += m * idx
			m *= n
		}
		return res
	} else {
		return 0
	}
}

func find(base string, x rune) int {
	for i, c := range base {
		if c == x {
			return i
		}
	}
	return -1
}
