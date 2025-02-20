package main

import "fmt"

func main() {
	// Aritiméticos
	soma := 1 + 2
	sub := 1 - 2
	div := 10 / 4
	mult := 10 * 5
	restoDaDiv := 10 % 2

	fmt.Println(soma, sub, div, mult, restoDaDiv)

	var num1 int =10
	var num2 int = 58
	soma2 := num1 + num2
	fmt.Println(soma2)

	//Operadores Relacionados
	fmt.Println( 1>2)
	fmt.Println( 1>=2)
	fmt.Println( 1==2)
	fmt.Println( 1<=2)
	fmt.Println( 1<2)
	fmt.Println( 1!=2)

	//Operadores Lógicos
	num :=10
	num++
	fmt.Println(num)

	num += 2
	fmt.Println(num)

	num --
	num *=2
	num/=2
	fmt.Println(num)

	num %=3
	fmt.Println(num)

}