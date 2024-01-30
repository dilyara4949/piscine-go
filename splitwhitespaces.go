package piscine

func SplitWhiteSpaces(s string) []string {
	var ans []string
	tmp := ""
	for _, c := range s {
		if c == ' ' {
			if tmp != "" {
				ans = append(ans, tmp)
			}
			tmp = ""
		} else {
			tmp += string(c)
		}
	}
	if tmp != "" {
		ans = append(ans, tmp)
	}
	return ans
}
