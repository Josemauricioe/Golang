package main

import (
	"fmt"
	"math"
	
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura, largura float64
}

type circulo struct {
	raio float64
}

func escreverArea(f forma) {
	fmt.Printf("\nA area da forma é %0.2f", f.area())
}

func (r retangulo) area() float64{
	return r.altura * r.largura
}

func (c circulo) area() float64{
	return math.Pi * (math.Pow(c.raio,2))
}

func main() {

	p := retangulo{10,15}
	escreverArea(p)
	c1 := circulo{10}
	escreverArea(c1)
}