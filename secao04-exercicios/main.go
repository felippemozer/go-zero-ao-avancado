package secao04exercicios

import "fmt"

// Ex 01: Criar um array com 2 posições de inteiros e armazenar em uma variável a soma total da lista. A variável deve ser impressa no console.

func ex01() {
	fmt.Println("Seção 04 - Exercicio 01:")
	arr := [2]int{54, 89}
	var soma int
	for _, v := range arr {
		soma += v
	}
	fmt.Println(soma)
}

/*
Ex 02: Dado um slice com os itens "2,8,3,10,5,4,7,9,1" que vão de 1 a 10,
	efetuar soma em duas variáveis, a primeira números de 1 a 5 e a segunda de 6 a 10.
	Imprimir os dois resultados;
*/

func ex02() {
	fmt.Println("Seção 04 - Exercicio 02:")
	s := []int{2, 8, 3, 10, 5, 4, 4, 7, 9, 1}
	var somaMenores, somaMaiores int

	for _, v := range s {
		if v <= 5 {
			somaMenores += v
		} else {
			somaMaiores += v
		}
	}

	fmt.Println("Menores:", somaMenores)
	fmt.Println("Maiores:", somaMaiores)
}

func Main() {
	ex01()
	ex02()
}
