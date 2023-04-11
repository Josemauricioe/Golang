package main

import "fmt"

func calculoMatematicos(n1, n2 int) (soma, subtração int) {
	soma = n1 + n2
	subtração = n1 - n2
	return
}

func main() {

	soma, subtracao := calculoMatematicos(20, 10)
	fmt.Println(soma,subtracao)

}