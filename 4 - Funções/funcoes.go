package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func variaveis(){
	var f = func (txt string)  string{
		fmt.Println(txt)
		return txt
	}
	var x = func(){
		fmt.Println("Função X")
	}
	x()
	f("Texto da Função 1")
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)
}
