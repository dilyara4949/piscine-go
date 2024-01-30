package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	sign := 0
	m := 1
	if s[0] == '-' {
		sign = -1
	} else if s[0] == '+' {
		sign = 1
	}
	end := 0
	if sign != 0 {
		end = 1
	}
	for i := len(s) - 1; i >= end; i-- {
		if rune(s[i])-'0' > 9 || rune(s[i])-'0' < 0 {
			return 0
		}
		res += m * int(rune(s[i])-'0')
		m *= 10
	}
	if sign == 0 {
		sign = 1
	}
	return res * sign
}
