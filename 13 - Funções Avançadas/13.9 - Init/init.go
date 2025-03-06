package main

import "fmt"

func main() {
	fmt.Println("Função main sunedo executada")
}

// Pode ser usada por arquivo, diferente da função main que funciona por pacote
func init() {
	fmt.Println("Executando função init")
}