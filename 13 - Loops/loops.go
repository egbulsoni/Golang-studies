package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println("Incrementando j")
	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j")
	// 	time.Sleep(time.Second)
	// }

	// nomes := [3]string{"Joao", "Davi", "Lucas"}

	// for indice, nome := range nomes{
	// 	fmt.Println(indice, nome)
	// }

	// for indice, letra := range "PALAVRA"{
	// 	fmt.Println(indice, string(letra))
	// }

	usuario := map[string]string {
		"nome": "Leonardo",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// nao da pra usar range em structs!!
	
	// type usuarioStruct struct {
	// 	nome      string
	// 	sobrenome string
	// }

	// usuario2 := usuarioStruct{"Ze", "Junior"}
	
	// for chave, valor := range usuario2 {
	// 	fmt.Println(chave, valor)
	// }
}