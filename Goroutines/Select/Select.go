package main

import (
	"fmt"
	"time"
)

func main() {

	canal, canal1 := make(chan string), make(chan string)

	go func() {
		for{
		time.Sleep(time.Millisecond * 500)
		canal <- "Canal"
		}
	}()
	go func() {
		for{
		time.Sleep(time.Millisecond * 500)
		canal1 <- "Canal 1"
		}
	}()

	for{

		select{
		case mensagemCanal := <- canal:
			fmt.Println(mensagemCanal)
		case mensagemCanal1 := <- canal1:
			fmt.Println(mensagemCanal1)
		}

		// mensagemCanal := <- canal
		// fmt.Println(mensagemCanal)

		// mensagemCanal2 := <- canal1
		// fmt.Println(mensagemCanal2)
	}
}