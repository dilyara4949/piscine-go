package piscine

func Unmatch(a []int) int {
	f := make([]bool, len(a))
	for i, c := range a {
		if !f[i] {
			f[i] = true
			ok := false
			for j := i + 1; j < len(a); j++ {
				if !f[j] && c == a[j] {
					f[j] = true
					ok = true
					break
				}
			}
			if !ok {
				return c
			}
		}
	}
	return -1
}
