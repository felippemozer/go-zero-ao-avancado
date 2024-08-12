package interfacevazia

import "fmt"

func Main() {
	// Ter múltiplos tipos para a mesma variável
	var a interface{}
	a = 5
	fmt.Println(a)

	a = circulo{
		raio: 10,
	}
	fmt.Println(a)

	// Ex: Slices de vários tipos
	var lista []interface{}
	lista = append(lista, 10)
	lista = append(lista, true)
	lista = append(lista, "teste")
	lista = append(lista, 525)

	// Verificação de tipo
	for _, v := range lista {
		if _, ok := v.(string); ok {
			fmt.Println(v, "string")
		} else {
			fmt.Println(v, "outros")
		}
	}
}
