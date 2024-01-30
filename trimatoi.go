package piscine

func TrimAtoi(s string) int {
	res := 0
	sign := 1
	idx := -1
	m := 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			sign = -1
			idx = i
		}
		if s[i] >= '0' && s[i] <= '9' {
			res += int(rune(s[i])-'0') * m
			m *= 10
			if i < idx {
				sign = 1
			}
		}
	}
	return res * sign
}
