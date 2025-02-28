package main

import "fmt"

func main() {

	// Criação de uma array, começando por 0 e indo até o valor 2
	var array1 [3] int8
	fmt.Println(array1)
	
	// Array de Strings
	array2 := [3]string {"New array","Second value", "Third value"}
	fmt.Println(array2)

	// Array de boolean, lembrando que a ordem dos fatores altera o produto
	array3 := [3] bool {false, true , false}
	fmt.Println(array3)

	// Numero de linhas é infinita nesse tipo de array, aumenta de acordo com o numero de parâmetros definidos na mesma linha
	arrayInf := [...]int {1,2,3,4,5}
	fmt.Println(arrayInf)

	//Slice tem sua capacidade infinita de linhas por padrão
	slice:= []int{1,2,3,5,7,11,13,17,19,23} 
	fmt.Println(slice)

	// Posso adicionar valores a slices ja criadas previamentes com essa formula
	slice = append(slice, 29)
	fmt.Println(slice)

	//Slice pode ser definido como uma parte do Array, com eles também podemos retirar parte do parâmetros, como o caso abaixo:
	// 1:3 significa que ele quer pegar o que está entre, sendo eles o índice 1 (inclusivo) e até o 2 (exclusivo), lembrando que começa apartir do 0
	slice2 := array2 [1:3]
	fmt.Println(slice2)
	//Nesse caso, diminuimos o tamanho do Array e o Valor 0 ("New Array") foi removido 

	//Nesse caso substituiremos o segundo valor, e como retiramos o 0 previamente, ele se torna o primeiro na ordem crescente 
	array2[1] = "New Array"
	fmt.Println(array2)

	//Arrays Internos
	fmt.Println("---------------------------")
	// Quando não especificado a criação do Array, ele cria um Array Interno armazenado na memória (??? a confirmar)
	slice3 :=make([]float32,10 ,11)
	fmt.Println(len(slice3)) // Tamanho da linha
	fmt.Println(cap(slice3)) // Capacidade Total
	fmt.Println("---------------------------")
	//Arrays internas aumentam dinamicamente de acordo com a sua demanda, ou seja , 
	// podem aumentar seu tamanho limite mesmo ja previamente tendo sua capacidade estabelecida
	slice3 = append(slice3, 5,6,7,8,9)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) 
	fmt.Println(cap(slice3)) 

	slice4 := make([]float32,4)
	fmt.Println(slice4)
	slice4 =append(slice4, 8)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}