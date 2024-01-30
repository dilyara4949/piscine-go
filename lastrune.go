package piscine

func LastRune(s string) rune {
	t := []rune(s)
	return t[len(t)-1]
}
