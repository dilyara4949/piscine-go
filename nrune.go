package piscine

func NRune(s string, n int) rune {
	t := []rune(s)
	if len(t) >= n && n >= 1 {
		return t[n-1]
	}
	return 0
}
