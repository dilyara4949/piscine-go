package piscine

func IterativeFactorial(nb int) int {
	if nb > 25 || nb < 0 {
		return 0
	}
	res := 1
	for i := 2; i <= nb; i++ {
		res *= i
	}
	return res
}
