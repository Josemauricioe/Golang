package main

import (
	"fmt"
	"time"

	
)

func main() {

	canal := multiplexar(escrever("Ola mundo"), escrever("Progrmando em agora"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canaldeEntrada1, canaldeEntrada2 <- chan string) <- chan string{
	canalDesaida := make(chan string)

	go func ()  {
		for{
			select{
			case messagem := <- canaldeEntrada1:
				canalDesaida <- messagem
			case messagem := <- canaldeEntrada2:
				canalDesaida <- messagem
			}
		}
	}()
	return canalDesaida
}

func escrever(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebindo: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}