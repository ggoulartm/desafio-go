package main

import "fmt"

func main() {
	x := 0
	for x < 100 {
		x++
		if x%3 != 0 {
			continue
		}
		fmt.Printf("%d é divisível por 3 \n", x)
	}
}
