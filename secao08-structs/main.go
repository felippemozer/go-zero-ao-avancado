package secao08structs

import (
	"fmt"
	"time"
)

func Main() {
	e := exEndereco()
	exPessoa(e)
	exAutomovel()
}

func exEndereco() endereco {
	fmt.Println("Seção 08 - Endereço:")
	e := endereco{
		cidade: "Vitória",
		numero: 35,
		rua:    "Av. Nossa Senhora da Penha",
	}
	fmt.Println(e)
	e.numero = 18
	fmt.Println(e.numero)
	return e
}

func exPessoa(e endereco) {
	fmt.Println("Seção 08 - Pessoa:")
	p := pessoa{
		nome:           "João",
		endereco:       e,
		dataNascimento: time.Date(2001, time.December, 31, 0, 0, 0, 0, time.Local),
	}
	p.calculaIdadeESalva()
	fmt.Println(p)
}

func (p *pessoa) calculaIdadeESalva() {
	anoNascimento := p.dataNascimento.Year()
	anoAtual := time.Now().Year()

	p.idade = anoAtual - anoNascimento
}

func exAutomovel() {
	fmt.Println("Seção 08 - Herança:")

	m := moto{
		automovel: automovel{
			ano:    2024,
			placa:  "XPTO-123",
			modelo: "CG",
		},
		cilindradas: 150,
	}
	fmt.Println("Moto:", m.modelo)

	c := carro{
		automovel: automovel{
			ano:    2024,
			placa:  "XPTO-123",
			modelo: "Fiesta",
		},
		portas:               4,
		potencia:             1600,
		possuiArCondicionado: false,
	}
	fmt.Println("Carro:", c.modelo)
}
