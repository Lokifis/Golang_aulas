package main

import "fmt"

func invertSinal(num int) int {
	return num * -1
}
func invertSinalComPont(num *int) {
	*num = *num * -1 	
}

func main() {
	num := 20
	numInvert := invertSinal(num)
	fmt.Println(numInvert)
	fmt.Println(num)

	novoNum:=40
	fmt.Println(novoNum)
	invertSinalComPont(&novoNum)
	fmt.Println(novoNum)
}