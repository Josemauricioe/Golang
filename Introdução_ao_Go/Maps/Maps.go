package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	usuario := map[string]string{
		"Nome": "Pedro",
		"Sobrenome": "Silva",
	}
	fmt.Println(usuario["Nome"])

	u2 := map[string]map[string]string{
		"nome": {
			"primeiro" : "Jose",
			"Segundo" : "João",
		},
		"Curso": {
			"nome" : "T.I",
			"Ano" : "Teiceiro",
		},
	}
	fmt.Println(u2)
	
	
}