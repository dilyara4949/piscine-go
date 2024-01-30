package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for _, c := range arg[1:] {
		for _, k := range c {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}
