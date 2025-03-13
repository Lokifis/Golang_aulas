package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Ola Mundo!"
	canal <- "Programando em Go"

	mensagem := <-canal
	mensagem2 := <-canal
	// Caso ultrapasse a capacidade, cai em Deadlock
	//mensagem3 := <- canal

	fmt.Println(mensagem, mensagem2)
}