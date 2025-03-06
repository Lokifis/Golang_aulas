package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}
func main (){
	texto := "Dentro da Main"
	fmt.Println(texto)

	//Função Closure armazena informações da função referenciada anteriormente 
	funcaoNova := closure()
	funcaoNova()
}