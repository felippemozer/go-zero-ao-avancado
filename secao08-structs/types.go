package secao08structs

import "time"

type endereco struct {
	rua    string
	numero int
	cidade string
}

type pessoa struct {
	nome           string
	endereco       endereco
	dataNascimento time.Time
	idade          int
}

type automovel struct {
	ano    int
	placa  string
	modelo string
}

type moto struct {
	automovel
	cilindradas int
}

type carro struct {
	automovel
	portas               int
	potencia             int
	possuiArCondicionado bool
}
