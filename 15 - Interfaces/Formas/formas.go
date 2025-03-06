package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func escreverArea(f forma) {
	fmt.Println("Área da forma é &0.2f", f.area())
}

// Nesse caso, fazemos com que os requisitos sejam cumpridos de acordo com sua condição
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// O método precisa ser do mesmo tipo se não ele não pode ser considerado um
func main() {
	r := retangulo{10, 15}
	escreverArea(r)
}
