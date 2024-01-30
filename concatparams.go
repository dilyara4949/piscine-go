package piscine

func ConcatParams(args []string) string {
	ans := ""
	for i, s := range args {
		ans += s
		if i < len(args)-1 {
			ans += "\n"
		}
	}
	return ans
}
