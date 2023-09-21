package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64
	// var numero int8 = 10000 // Overflow
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// int 32 = rune
	var numero3 rune = 12456
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)


	// float32, float64
	var numeroReal1 float32	= 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// var str string = "Texto"
	// var booleano1 bool = true

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}