package main

import "fmt"

func main() {
	// func(texto string) {
	// 	fmt.Println(texto)
	// }("Passando Parametro")
	retorno := func(texto string) string{
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parametro")

	fmt.Printf(retorno)
}