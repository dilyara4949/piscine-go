package piscine

func Compact(s *[]string) int {
	ans := []string{}
	for i := range *s {
		if (*s)[i] != "" {
			ans = append(ans, (*s)[i])
		}
	}
	*s = ans
	return len(ans)
}
