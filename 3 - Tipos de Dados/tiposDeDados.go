package main

import "fmt"

func main() {
	// Números inteiros | int8 até 64 , uint8 até 64 e sem especificação
	// Valor do int baseado na arquitetura do pc
	numero := 1000000000000000000
	fmt.Println("O número é: ", numero)

	// Unsigned int = aceita apenas valores positivos
	var numero1 uint32 = 1000000000
	fmt.Println("O número1 é: ", numero1)

	// Inteiro de 32 bits
	var numero2 int32 = -1000000000
	fmt.Println("O número2 é: ", numero2)
	// ---------------------------
	// Números Reais | float32 e float64, sem espec.
	numero3 := 198734982.24141
	fmt.Println("O número3 é: ", numero3)

	// ---------------------------
	var numero4 string = "10000000000"
	fmt.Println("O número4 é: ", numero4)

	var numero5 bool = true
	fmt.Println("O número5 é: ", numero5)

	// Alias | byte e rune
	// Byte = uint8
	var numero6 byte = 100
	fmt.Println("O número6 é: ", numero6)

	// rune = int32
	var numero7 rune = 'a'
	fmt.Println("O número7 é: ", numero7)

	//Pega o número diretamente da tabela ASCII
	char := 'B'
	fmt.Println(char)
}
