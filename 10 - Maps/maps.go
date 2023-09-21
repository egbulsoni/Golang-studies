package main

import "fmt"

func main(){
	fmt.Println("Maps")
	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "Joao",
			"ultimo": "Pedro",
		},
		"curso": {
			"nome": "engenharia",
			"campus": "campus 1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string {
		"nome": "gemeos",
	}

	fmt.Println(usuario2)
}