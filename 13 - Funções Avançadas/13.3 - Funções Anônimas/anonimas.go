package main

import "fmt"

//Função onde os parâmetros são implícitos, que deixa ela anônima
func main() {
	func(texto string) {
		fmt.Println(texto)
	}("Hello World!")

	retorno := func(texto2 string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto2, 13)
	}("Hello World 2")

	fmt.Println(retorno)
}
