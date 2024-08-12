package secao09exercicios

import (
	"fmt"
	"golang_zero_ao_avancado/secao09-exercicios/model"
	"time"
)

/*
- Criar um modelo que irá receber os itens para compra do mês, nesse modelo teremos a data que a compra irá acontecer, mercado e os itens para comprar.
- Dado o exercício anterior, mover o modelo anterior criado para o pacote chamado model
- Dado o exercício anterior, criar uma função no pacote model que inicializa a struct e retorno como ponteiro
*/

func Main() {
	arroz := model.Item{
		Nome:       "arroz",
		Quantidade: 1,
	}
	feijao := model.Item{
		Nome:       "feijão",
		Quantidade: 2,
	}
	oleo := model.Item{
		Nome:       "óleo",
		Quantidade: 3,
	}

	extrabom := model.Mercado{
		Nome: "extrabom",
	}

	lista := model.Compra{
		DataCompra: time.Now(),
		Mercado:    &extrabom,
		Itens: []*model.Item{
			&arroz,
			&feijao,
			&oleo,
		},
	}
	fmt.Println(lista.DataCompra)
	fmt.Println(lista.Mercado.Nome)
	for _, v := range lista.Itens {
		fmt.Println(v.Nome, v.Quantidade)
	}
}
