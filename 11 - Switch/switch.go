package main

import "fmt"

//Switchs
func diaDaSemana(num int) string {
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Inválido"
	}
}
// Segundo Exemplo
func diaDaSemana2(num int) string {
	var diaDaSemana string
// Nesse exemplo, declaramos o valor da variavel específico 
	switch{
	case num ==1:
		diaDaSemana="Domingo"
		fallthrough
	case num ==2:
		diaDaSemana="Segunda"
	case num==3:
		diaDaSemana="Terça-feira"
	case num==4:
		diaDaSemana="Quarta-feira"
	case num==5:
		diaDaSemana="Quinta-feira"
	case num==6:
		diaDaSemana="Sexta-feira"
	case num==7:
		diaDaSemana="Sábado"
	default:
		diaDaSemana="Inválido"
	}	
	return diaDaSemana
}
	func main() {

		// Retorna valor de acordo com o parametro especificado
		dia:= diaDaSemana(3)
		fmt.Println(dia)

		dia2:= diaDaSemana(8)
		fmt.Println(dia2)

		fmt.Println("-----------------")

		// Nesse exemplo, o Fallthrough joga o valor para dentro da próxima função 
		dia3:= diaDaSemana2(1)
		fmt.Println(dia3)
}