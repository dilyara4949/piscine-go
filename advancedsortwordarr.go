package piscine

func AdvancedSortWordArr(t []string, f func(a, b string) int) {
	n := len(t)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if f(t[j], t[k]) > 0 {
					tmp := t[j]
					t[j] = t[k]
					t[k] = tmp
				}
			}
		}
	}
}
