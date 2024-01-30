package piscine

func IsPrintable(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !(c >= 32 && c <= 126) {
			return false
		}
	}
	return true
}
