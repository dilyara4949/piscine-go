package piscine

func ActiveBits(n int) int {
	cnt := 0
	for n != 0 {
		if n%2 == 1 {
			cnt++
		}
		n /= 2
	}
	return cnt
}
