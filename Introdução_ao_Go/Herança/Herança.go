package main

import "fmt"

type pessoa struct {
	nome, sobrenome string
	idade           int8
}

type estudande struct {
	pessoa
	curso string
	ano   int8
}

func main() {

	fmt.Println("Herança")

	p1 := pessoa{"Jose","Maurício",23}
	fmt.Println(p1)

	e1 := estudande{p1,"INFORMATICA",3}
	fmt.Println(e1.nome)

}