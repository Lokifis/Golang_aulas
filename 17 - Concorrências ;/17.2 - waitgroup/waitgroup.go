package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// waitgroups são funcionalidades que podem ser especificadas a quantidade de goroutines com a finalidade de sincronizar elas. EX:
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func ()  {
		escrever("Olá Mundo!")
		waitGroup.Done()
	}()
	go func ()  {
		escrever("Programando em Go!")
		waitGroup.Done()	
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i:= 0; i<13; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
