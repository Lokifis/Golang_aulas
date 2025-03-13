package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrência X Paralelismo

	// Paralelismo ocorre somente quando 2 ou mais tarefas são executadas ao mesmo tempo (somente possível em processadores com mais de 1 núcleo)
	// Concorrência pode ocorrer enquanto 1 ou mais tarefas estão sendo executadas, ou seja, elas não necessariamente depende do término da tarefa executada (possível em processador com 1 núcleo)

	go escrever("Olá Mundo!") // goroutine = independente da função ser infinita, execute a próxima
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
