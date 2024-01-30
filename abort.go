package piscine

func Abort(a, b, c, d, e int) int {
	s := []int{a, b, c, d, e}
	SortIntegerTable(s)
	return s[2]
}
