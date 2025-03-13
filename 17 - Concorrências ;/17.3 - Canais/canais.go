package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	fmt.Println("Oi Mundo!")

	// for {
	// 	mensagem, aberto := <-canal
	// 	if aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	// Código acima tem a mesma funcionalidade que o debaixo
	
	for mensagem := range canal {

		fmt.Println(mensagem)
	}
	fmt.Println("Fim.")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
		// Caso não tenha nada para ser seguido após, a um risco de cair na Deadlock (Erro de execução, visivel apenas na compilação) 
	}

	close(canal)
}