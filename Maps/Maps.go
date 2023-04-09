package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	usuario := map[string]string{
		"Nome": "Pedro",
		"Sobrenome": "Silva",
	}
	fmt.Println(usuario["Nome"])
}