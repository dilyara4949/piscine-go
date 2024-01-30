package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	for _, s := range arg {
		if s == "01" || s == "galaxy" || s == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
