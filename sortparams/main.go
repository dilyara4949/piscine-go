package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := 1; i < len(arg); i++ {
		for j := 1; j < len(arg); j++ {
			for k := j + 1; k < len(arg); k++ {
				if arg[j] > arg[k] {
					tmp := arg[j]
					arg[j] = arg[k]
					arg[k] = tmp
				}
			}
		}
	}
	for _, c := range arg[1:] {
		for _, k := range c {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}
