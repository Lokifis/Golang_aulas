package main

import "fmt"

func fibbonacci(posicao uint) uint {
	//Cuidado! pode dar Stack Overflow
	if posicao <= 1 {
		return posicao
	}
	return fibbonacci((posicao - 2) + fibbonacci(posicao-1))
}

func main() {
	fmt.Println("Sequência de Fibbonacci")
	posicao := uint(10)
	fmt.Println(fibbonacci(posicao))
}
