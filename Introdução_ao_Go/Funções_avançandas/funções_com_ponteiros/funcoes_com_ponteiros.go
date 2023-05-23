package main

import (
	"fmt"
	
)

func inverterSinal(n1 int) int {
	return n1 * -1
}
func inverterSinalcomPonteiros(n1 *int) {
	*n1 = *n1 * -1 
}
func main() {
	numero := 20
	numeroIntertido := inverterSinal(numero)
	fmt.Println(numeroIntertido)
	fmt.Println(numero)

	novo_n1 := 40
	fmt.Println(novo_n1)
	inverterSinalcomPonteiros(&novo_n1)
	fmt.Println(novo_n1)
}