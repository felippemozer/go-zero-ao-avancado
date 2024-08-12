package cominterfaces

import (
	"fmt"
	"math"
)

func Main() {
	fmt.Println("Com interfaces: ")
	r := retangulo{
		largura: 5,
		altura:  4,
	}

	c := circulo{
		raio: 12,
	}

	exibirArea(r)
	exibirArea(c)
}

func (r retangulo) area() float64 {
	fmt.Printf("Área retângulo: ")
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	fmt.Printf("Área círculo: ")
	return math.Pi * math.Pow(c.raio, 2)
}

func exibirArea(g geometria) {
	fmt.Println(g.area())
}
