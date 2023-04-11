package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros{
		total += numero
	}
	return total
}

func escrever(text string,numeros ...int){
	for _, numeros := range numeros{
		fmt.Println(numeros)
	}
}

func main() {

	totalsoma := soma(1,2,3,4)
	fmt.Println(totalsoma)

	escrever("Ola mundo", 10,2,3,4,5,6,7)

}