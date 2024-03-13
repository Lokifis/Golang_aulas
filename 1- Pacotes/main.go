package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

// Function shows messages from main and auxiliar
func main() {
	fmt.Println("Hello, world!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}
