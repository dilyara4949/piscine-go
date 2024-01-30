package piscine

func Map(f func(int) bool, a []int) []bool {
	ans := []bool{}
	for _, i := range a {
		ans = append(ans, f(i))
	}
	return ans
}
