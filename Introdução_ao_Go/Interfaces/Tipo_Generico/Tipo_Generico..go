package main

import "fmt"

func generica(intef interface{}) {
	fmt.Println(intef)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	fmt.Println(1,2, "String", false, true, float64(14252))
	
}