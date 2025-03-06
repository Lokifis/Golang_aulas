package main

import "fmt"

func recuperarExec(){
	if r:= recover(); r!= nil{
		fmt.Println("Tentativa de recuperação")
	}
}

func whitelist(numId, twoSteps int32) bool {
	defer recuperarExec()
	id := (numId + twoSteps)

	if id > 0 {
		return true
	} else if id < 0 {
		return false
	}
	// antes de dar pânico, a linguagem Go vai tentar executar as funções adiadas pelo Defer 
	panic("Inexistente!")
}
func main() {
	fmt.Println(whitelist(1,1))
}
