package piscine

func CollatzCountdown(start int) int {
	if start < 1 {
		return -1
	}
	cnt := 0
	for start != 1 {
		if start%2 == 0 {
			start /= 2
		} else {
			start = start*3 + 1
		}
		cnt++
	}
	return cnt
}
