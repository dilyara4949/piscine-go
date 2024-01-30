package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}
	ans := make([]int, max-min)
	for i := min; i < max; i++ {
		ans[i-min] = i
	}
	return ans
}
