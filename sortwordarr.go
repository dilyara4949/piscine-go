package piscine

func SortWordArr(t []string) {
	n := len(t)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if t[j] > t[k] {
					tmp := t[j]
					t[j] = t[k]
					t[k] = tmp
				}
			}
		}
	}
}
