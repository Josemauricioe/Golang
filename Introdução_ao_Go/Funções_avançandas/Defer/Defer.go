package main

import "fmt"

func fum1() {
	fmt.Println("Print fun1")
}
func fum2() {
	fmt.Println("Print fun2")
}

func notas(n1,n2 float64)bool{
	defer fmt.Println("Media calculada. Resultado sera retornado")
	fmt.Println("Entrando na funçao para verificar se o aluno foi aprovando")
	media := (n1+n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {	
		
	defer fmt.Println(notas(7,8))
	defer fum1()
	fum2()

}