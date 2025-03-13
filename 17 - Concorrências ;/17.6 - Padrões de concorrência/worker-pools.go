package main

import "fmt"

func main() {
	tarefas := make(chan int, 32)
	resultados := make(chan int, 32)

	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for num := range tarefas {
		resultados <- fibbonacci(num)
	}
}

func fibbonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibbonacci((posicao - 2) + fibbonacci(posicao-1))
}
