package piscine

func Rot14(s string) string {
	ans := ""
	for _, c := range s {
		k := c
		if c >= 'a' && c <= 'z' {
			k = c + 14
			if k > 'z' {
				k = 'a' + (13 - ('z' - c))
			}
		} else if c >= 'A' && c <= 'Z' {
			k = c + 14
			if k > 'Z' {
				k = 'A' + (13 - ('Z' - c))
			}
		}
		ans += string(k)
	}
	return ans
}
