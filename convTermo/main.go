package main

import "fmt"

func main() {
	fmt.Println("O desafio é converter a temperatura de ebulição da água de Kelvin para Celsius")
	tempKelvinDeEbulicao := 373
	fmt.Printf("A temperatura de ebulição da água em Kelvin é %d \n", tempKelvinDeEbulicao)
	tempCelsiusDeEbulicao := tempKelvinDeEbulicao - 273
	fmt.Println("A fórmula de conversão de unidades termométricas de Celsius para Kelvin é Celsius=Kelvin-273")
	fmt.Printf("Portanto, a temperatura de ebulição da água em Celsius é %d \n", tempCelsiusDeEbulicao)
	fmt.Println(tempCelsiusDeEbulicao)
}
