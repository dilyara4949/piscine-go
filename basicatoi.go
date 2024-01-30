package piscine

func BasicAtoi(s string) int {
	res := 0
	m := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += m * int(rune(s[i])-'0')
		m *= 10
	}
	return res
}
