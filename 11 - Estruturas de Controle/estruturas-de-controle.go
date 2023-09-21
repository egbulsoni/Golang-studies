package main

import "fmt"

func main(){
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// if init morre nesse escopo
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero eh maior q zero")
	} else {
		fmt.Println("Numero eh menor que zero")
	}

	// fmt.Println(outroNumero)
}