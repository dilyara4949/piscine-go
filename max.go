package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	mx := a[0]
	for _, i := range a {
		if i > mx {
			mx = i
		}
	}
	return mx
}
