package main

import "fmt"

func calcalado(n1, n2 int8) (int8,int8){
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao;
}

func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	soma := soma(20, 20)
	fmt.Println("A soma de n1 e n2 é: ",soma)

	var f = func (txt string) string  {
		fmt.Println(txt)
		return txt
	}

	resultado := f("TEXTO DA FUNÇÂO RESULTADO")
	fmt.Println(resultado)
		
	resultado_calculos_soma,resultado_calculos_subtracao := calcalado(32,42)
	fmt.Println(resultado_calculos_soma, resultado_calculos_subtracao)
	
	//resultado_calculos_soma,  _ := calcalado(32,42)
	//fmt.Println(resultado_calculos_soma)
}
