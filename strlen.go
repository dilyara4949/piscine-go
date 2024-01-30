package piscine

func StrLen(s string) int {
	cnt := 0
	for _, i := range s {
		cnt++
		i++
	}
	return cnt
}
