package main

import "fmt"

// Funções recursivas são funções que chamam a si mesmas
// Fibbonacci é a soma da sequencia dos parâmetros, exemplo:
// 1 -> 1 -> 2 -> 3 -> 5 -> 8 -> 13

func fibbonacci(posicao uint) uint {
	//Cuidado! pode dar Stack Overflow
	if posicao <= 1 {
		return posicao
	}
	return fibbonacci((posicao - 2) + fibbonacci(posicao-1))
}

func main() {
	fmt.Println("Sequência de Fibbonacci")
	posicao := uint(5)

	for i:= uint (0); i< posicao; i++{
		fmt.Println(fibbonacci(i))
	}
}
