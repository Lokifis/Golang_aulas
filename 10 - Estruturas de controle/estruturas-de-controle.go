package main

import "fmt"

func main() {
	num := 13

	//If e Else básico
	if num >= 15 {
		fmt.Println("Maior ou igual a 15")
	}else{
		fmt.Println("Menor que 15")
	}


	// Quando uma variável é criada dentro da estrutura, ela é limitada a somente aquela estrutura 
	if num2 := num; num2>0{
		fmt.Println("Maior que 0")
	}else if num == 0{
		fmt.Println("Igual a 0")
	}else{
		fmt.Println("Menor que 0")
	}
}