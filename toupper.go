package piscine

func ToUpper(s string) string {
	res := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			res += string(rune(c - 32))
		} else {
			res += string(c)
		}
	}
	return res
}
