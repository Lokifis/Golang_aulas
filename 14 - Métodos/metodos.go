package main

import "fmt"

//Métodos são obrigatoriamente dentro de uma estrutura, sem pontas soltas como as funções
type user struct {
	nome  string
	idade uint8
}

func (u user) salvar() {
	fmt.Println("Salvo!")
}

func (u user) maiorDeIdade() bool {
	return u.idade >= 18
}

//Ponteiros utilizamos para atualizar valores
func (u *user) aniversario() {
	u.idade++
}

func main() {
	user1 := user{"Usuário 1", 13}
	user1.salvar()

	user2 := user{"Usuário 2", 26}
	user2.salvar()

	maiorDeIdade := user1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	user2.aniversario()
	fmt.Println(user2.idade)
}
