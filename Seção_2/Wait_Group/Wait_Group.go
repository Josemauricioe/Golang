package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wail_group sync.WaitGroup

	wail_group.Add(2)

	go func ()  {
		escrever("Ola mundo")
		wail_group.Done()
	}()
	go func ()  {
		escrever("Programando em Go")
		wail_group.Done()
	}()

	wail_group.Wait()	
}

func escrever(texto string) {
	for i := 0; i < 5 ;i++ {
		fmt.Println(texto)
		time.Sleep(time.Second);	
	}
}