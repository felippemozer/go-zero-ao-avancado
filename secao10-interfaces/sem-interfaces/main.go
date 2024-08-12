package seminterfaces

import (
	"fmt"
	"math"
)

func Main() {
	fmt.Println("Sem interfaces: ")
	r := retangulo{
		largura: 1,
		altura:  2,
	}

	c := circulo{
		raio: 5,
	}

	exibirAreaRetangulo(r)
	exibirAreaCirculo(c)
}

func exibirAreaRetangulo(r retangulo) {
	area := r.altura * r.largura
	fmt.Println("Área retângulo:", area)
}

func exibirAreaCirculo(c circulo) {
	area := math.Pi * math.Pow(c.raio, 2)
	fmt.Println("Área círculo:", area)
}
