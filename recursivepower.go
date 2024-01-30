package piscine

func RecursivePower(n int, power int) int {
	if power == 0 {
		return 1
	}
	if power == 1 {
		return n
	}
	if power < 0 {
		return 0
	}
	return n * RecursivePower(n, power-1)
}
