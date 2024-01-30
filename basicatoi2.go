package piscine

func BasicAtoi2(s string) int {
	res := 0
	m := 1
	for i := len(s) - 1; i >= 0; i-- {
		if rune(s[i])-'0' > 9 || rune(s[i])-'0' < 0 {
			return 0
		}
		res += m * int(rune(s[i])-'0')
		m *= 10
	}
	return res
}
