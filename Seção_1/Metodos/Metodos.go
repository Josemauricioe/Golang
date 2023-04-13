package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

func (u usuario) salvar() {
	fmt.Println("Salvado os dados do Usuario %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeidade() bool {
	return u.idade >= 18
}

func (u *usuario) aniversario(){
	u.idade++
}

func main() {

	usuario1 := usuario{"Usuario 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Usuario 2", 30}
	usuario2 .salvar()
	
	usuario2.aniversario()
	fmt.Println(usuario2.idade)
}