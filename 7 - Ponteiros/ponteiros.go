package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	// Ponteiros são referências de memória
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	// Aponta para onde a memória esta armazenada
	fmt.Println(variavel3, ponteiro)
	
	//Aponta o valor armazenado na memória
	fmt.Println(variavel3, *ponteiro)
}
