package main

import (
	"fmt"

	
)

func main() {
	fmt.Println("Ponteiros")

	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1,v2)

	v1++

	fmt.Println(v1,v2)

	//Ponteiros é uma referência de memoria

	var v3 int
	var ponteiro1 *int

	v3 = 100
	ponteiro1 =&v3

	fmt.Println(v3,*ponteiro1)
	v3++
	fmt.Println(v3,*ponteiro1)
}