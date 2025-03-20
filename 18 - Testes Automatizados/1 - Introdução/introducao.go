package main

import (
	"fmt"
	canalDiscord "introducao-testes/Discord"
)

func main() {

	tipoCanal := canalDiscord.TipoDoCanal("Voip")
	fmt.Println(tipoCanal)
}