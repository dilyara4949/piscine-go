package piscine

func IterativePower(n int, power int) int {
	if power == 0 {
		return 1
	}
	if power == 1 {
		return n
	}
	if power < 0 {
		return 0
	}
	return n * IterativePower(n, power-1)
}
