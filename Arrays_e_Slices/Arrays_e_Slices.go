package main

import (
	"fmt"

	
)

func main() {
	fmt.Println("Arrays e Slices")

	//Formas de declarar
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]int{1,2,3,4,5}
	fmt.Println(array2)

	array3 := [...]int{1,2,3,4,5,6,7,8}
	fmt.Println(array3)

	//Fim de arrays

	slices := []int{11,12,12,13,15,15,16,20}
	fmt.Println(slices)

	slices = append(slices,18)

	fmt.Println(slices)

	slices2 := array2[1:3]
	fmt.Println(slices2)

	array2[1] = 100
	fmt.Println(slices2)

	//Arrays Internos

	slices3 := make([]float64,10,15)
	fmt.Println(slices3)
	

}