package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
	endereco endereco
}
type endereco struct {
	logradoura  string
	numero int8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "José"
	u.idade = 30
	fmt.Println(u)

	enderecoex := endereco{"Rua 1",32}

	u2 := usuario{"Davi",32,enderecoex}
	fmt.Println(u2)

	u3 := usuario{nome: "Davi"}
	fmt.Println(u3)
}