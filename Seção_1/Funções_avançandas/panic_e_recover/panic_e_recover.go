package main

import "fmt"

func recuperar(){
	if f := recover(); f  != nil{
		fmt.Println("Recupendo")
	}
}

func aprovando_OU_Reprovando(n1, n2 float64) bool {
	defer recuperar()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6{
		return false
	}
		panic("A MEDIA É EXATAMMENTE 6!!!")
}

func main() {	

	fmt.Println(aprovando_OU_Reprovando(7,6))
	fmt.Println("pos execução!")

}