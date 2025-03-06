package main

import "fmt"

func func1() {
	fmt.Println("Função 1")
}
func func2() {
	fmt.Println("Função 2")
}

func main (){
	// Defer serve para atrasar uma função, adiando quando ela será usada (dependendo de sua condicional)
	defer func1()
	func2()
}