package main

import (
	"fmt"
)

func main() {
	soma := sum(5, 10, 16)
	subtrai := sub(10, 4)
	multiplica := mult(10, 20, 30)
	divid := divider(500, 250)
	fmt.Println(soma, subtrai, multiplica, divid)
}

func sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum = sum + v
	}
	return sum
}

func sub(i ...int) int {
	sub := 0
	for _, v := range i {
		sub = v - sub
	}
	sub = -sub
	return sub
}

func mult(i ...int) int {
	mult := 1
	for _, v := range i {
		mult = mult * v
	}
	return mult
}

func divider(i ...int) int {
	div := 1
	for ind, v := range i {
		if ind == 0 {
			div = v
			continue
		}
		div = div / v
	}
	return div
}
