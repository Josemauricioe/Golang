package main

import (
	"fmt"
	
)

func main() {
	
	// func(text string) {
	// 	fmt.Println(text)
	// }("Passando Parâmetro")

	retoeno := func(text string)string {
		return fmt.Sprintf("Recebido -> %s", text)
	}("Passando Parâmetro")

	fmt.Println(retoeno)
	
}