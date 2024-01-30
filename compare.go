package piscine

func Compare(a, b string) int {
	i := 0
	mn := 0
	if len(a) <= len(b) {
		mn = len(a)
	} else {
		mn = len(b)
	}
	for ; i < mn; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	if len(a) > len(b) {
		return 1
	} else if len(a) < len(b) {
		return -1
	}
	return 0
}
