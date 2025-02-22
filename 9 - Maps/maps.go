package main

import "fmt"

func main() {
	
	// Maps aceitam somente do mesmo tipo
	user := map[string]string{
		"nome":      "Pedro",
		"Sobrenome": "Ribeiro",
	}
	fmt.Println(user)

	// A classificação deles é hierárquica, então existe essa possibilidade de "empilhar" mas não é muito recomendovel escalar muito
	user2:= map[string]map[string]string{
		"nome": {
			"primeiro":"Rainkily",
			"ultimo":"da Silva",
		},
		"Curso":{
			"nome": "Ciencias Exatas",
			"unidade":"UFV",
		},
	}	
	fmt.Println(user2)
	delete(user2, "nome")
	fmt.Println(user2)
	
	// Podemos adicionar também a Maps ja pré-estabelecidos
	user2["semestre"] = map[string]string{
		"unidade":"Primeiro",
	}
	fmt.Println(user2)
}