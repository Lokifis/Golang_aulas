package main

import "fmt"

type user struct {
	nome  string
	idade uint8
	email string
	identidade identidade
}

type identidade struct{
	cpf string
	rg string
} 

func main() {
	var u user
	u.nome = "Trevor"
	u.idade = 51
	u.email = "trevorSo@hotmail.com"
	fmt.Println(u)

	identidade := identidade{"99999","9999"}
	user2 := user{"Leo", 19, "leogamer@gmail.com", identidade}
	fmt.Println(user2)

	user3 := user{nome: "Jorge"}
	fmt.Println(user3)
}