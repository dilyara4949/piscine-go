package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for m := '0'; m <= '9'; m++ {
				for n := '0'; n <= '9'; n++ {
					if i == m && j < n || i < m {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(m)
						z01.PrintRune(n)
						if !(i == '9' && j == '8' && m == '9' && n == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
