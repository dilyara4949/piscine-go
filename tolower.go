package piscine

func ToLower(s string) string {
	res := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			res += string(rune(c + 32))
		} else {
			res += string(c)
		}
	}
	return res
}
