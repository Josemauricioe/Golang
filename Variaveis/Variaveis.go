package main

import "fmt"

func main() {
	var variave1 string = "variavel 1"
	variave2 := "variavel_2"
	fmt.Println(variave1)
	fmt.Println(variave2)

	var (
		variave3 string = "veaws"
		variave4 string = "vebae"
	)
	fmt.Print(variave3, variave4)
	
	variave5, vvariave6 := "Variaveis 5" , "Variaveis6"
	fmt.Println(variave5 , vvariave6)

	const variave1tel string = "constante 1"
	fmt.Print(variave1tel)

	variave5, vvariave6 = vvariave6 , variave5
	fmt.Println(variave5 , vvariave6)
	
}