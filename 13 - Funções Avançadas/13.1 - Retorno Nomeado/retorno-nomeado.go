package main

import "fmt"

// Função de soma simplificada e legível, que retornam os parametros dessa função sem precisar especificar
func calculosMatematicos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, sub := calculosMatematicos(13, 7)
	fmt.Println(soma, sub)
}
