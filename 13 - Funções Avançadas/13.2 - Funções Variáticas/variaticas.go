package main

import "fmt"

//Função de Soma que recebe os Inteiros informados e integra a um Slice
func soma(num ...int) int {
	total := 0
	for _, num := range num {
		total += num
	}
	return total
}

func escrever(texto string, num ...int) {
	for _, num := range num {
		fmt.Println(texto, num)
	}
}

func main() {
	totalDaSoma := soma(1, 2)
	totalDaSoma2 := soma(6, 4)
	totalDaSoma3 := soma(totalDaSoma2, totalDaSoma)

	escrever("Resultado da primeira soma:", totalDaSoma)
	escrever("Resultado da segunda soma:", totalDaSoma2)
	escrever("Resultado da soma entre ambas as funções:", totalDaSoma3)
}
