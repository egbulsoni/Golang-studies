package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}
func main(){
	fmt.Println("arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}
	u2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{idade: 32}
	fmt.Println(u3)
}