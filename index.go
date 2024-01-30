package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	for i, c := range s {
		if c == rune(toFind[0]) {
			if i+len(toFind) > len(s) {
				return -1
			}
			cnt := 0
			for j, l := range toFind {
				if l == rune(s[i+j]) {
					cnt++
				}
			}
			if cnt == len(toFind) {
				return i
			}
		}
	}
	return -1
}
