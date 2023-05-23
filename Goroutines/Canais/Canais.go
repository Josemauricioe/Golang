package main 

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string)
	canal1 := make(chan float64)

	go escrever("Ola mundo ", canal)
	go soma(23,42, canal1)
	
	fmt.Println("Depois da função escrever começar a ser executada")

	for mersagem := range canal{
		fmt.Println(mersagem)
	}
	// mersagem, aberto := <- canal	
	// 	if !aberto {
	// 	break
	// }
	// fmt.Println(mersagem)
	// }
	for{
		mersagem, aberto := <- canal1	
			if !aberto {
			break
		}
		fmt.Println(mersagem)
	}
	

	fmt.Println("Fim do Programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5 ;i++ {
		canal <- texto
		time.Sleep(time.Second);	
	}
	close(canal)
}

func soma(n1,n2 float64, canal chan float64) {
	for i := 0; i < 5 ;i++ {
		canal <- n1 + n2
		time.Sleep(time.Second);	
	}
	close(canal)
}