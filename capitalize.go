package piscine

func Capitalize(s string) string {
	res := []rune(s)
	if len(s) >= 1 && s[0] <= 'z' && s[0] >= 'a' {
		res[0] = res[0] - 32
	}

	for i, c := range res {
		if i > 0 && c <= 'z' && c >= 'a' && !IsAlpha(string(res[i-1])) {
			res[i] = res[i] - 32
		} else if i > 0 && c <= 'Z' && c >= 'A' && IsAlpha(string(res[i-1])) {
			res[i] = res[i] + 32
		}
	}
	ans := string(res)
	return ans
}
