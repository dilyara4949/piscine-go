package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("File name missing")
	} else if len(arg) > 1 {
		fmt.Println("Too many arguments")
	} else {
		file, err := os.ReadFile(arg[0])
		if err != nil {
			panic(err)
		}
		for i, c := range string(file) {
			if i < len(file)-1 {
				fmt.Print(string(c))
			}
		}
		fmt.Print("\n")
	}
}
