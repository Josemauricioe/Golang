package main

import (
	
	"fmt"
)

func main() {

	// newFunction()

	// for j := 0; j < 10; j++{
	// 	fmt.Println("Incrementado i")
	// 	time.Sleep(time.Second)
	// }

	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementado i")
	// 	time.Sleep(time.Second)
	// }


	// nomes:= [3]string{"Jose","Joao","Rafael"}

	// for _ , nomes := range nomes{
	// 	fmt.Println(nomes)
	// }

	for _ , letra := range "PALAVRA"{
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"Nome": "Pedro",
		"Sobrenome": "Silva",
	} 

	for chave, V := range usuario{
		fmt.Println(chave,V)
	}
}
