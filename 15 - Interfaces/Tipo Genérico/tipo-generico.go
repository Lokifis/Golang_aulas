package main

import (
	"fmt"
)

func generica(interf interface{}) {
	fmt.Println(interf)
}

// A função pode receber qualquer parâmetro independente de sua condição, que pode virar uma gambiarra caso não usada direito
func main() {
	generica("String")
	generica(1)
	generica(9.3)
	generica(true)

	fmt.Println(13, "Lokifis")

	// Isso é gambiarra k
	mapa := map[interface{}]interface{}{
		1:    "strings",
		true: float32(16),
	}
	fmt.Println(mapa)
}
