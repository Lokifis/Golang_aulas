package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()
	
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	// Extremamente parecido com Switch, sua diferença é que ele é usado majoritariamente na concorrência 
	for {
		select{
		case mensagemCanal1 := <- canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <- canal2:
			fmt.Println(mensagemCanal2)
		}
				// Se o resultado for cru, sem o select, o canal 1 esperaria até o canal 2 concluir, sendo que ele faria a compilação do texto muito mais rápido se não impedido
		// mensagemCanal1:= <- canal1
		// fmt.Println(mensagemCanal1)

		// mensagemCanal2:= <- canal1
		// fmt.Println(mensagemCanal2)
	}
}
