package main

import "fmt"

func main() {
	go escrever("Ola mundo") //goroutine
	escrever("Programa e Go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)	
	}
}