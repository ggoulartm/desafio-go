package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Desafio numérico - Parte 1")
	fmt.Println("Números divisíveis por 3 entre 1 e 100:")
	for x < 100 {
		x++
		if x%3 != 0 {
			continue
		}
		fmt.Printf("%d,\t", x)
	}
	fmt.Println("\n Desafio numérico - Parte 2")
	x = 0
	fmt.Println("\nQuando um número for divisível por 3, a mensagem será Pin")
	fmt.Println("Quando um número for divisível por 5, a mensagem será Pan\n ")
	for x < 100 {
		x++
		if x%3 == 0 {
			fmt.Printf("Pin\t")
		} else if x%5 == 0 {
			fmt.Printf("Pan\t")
		}
	}
}
