package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	n := IsPrime(nb)
	for !n {
		nb++
		n = IsPrime(nb)
	}
	return nb
}
