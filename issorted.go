package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 1; i < len(a); i++ {
		if f(a[0], a[len(a)-1]) < 0 {
			if f(a[i-1], a[i]) > 0 {
				return false
			}
		} else {
			if f(a[i-1], a[i]) < 0 {
				return false
			}
		}
	}
	return true
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}
