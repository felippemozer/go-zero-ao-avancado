package secao12recover

import (
	"fmt"
	"os"
)

func Main() {
	readFile()

	fmt.Println("Fim")
}

func readFile() {
	// Isso não é recomendado! Pode causar efeitos colaterais
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado")
		}
	}()

	_, err := os.Open("./settings.txt")
	if err != nil {
		panic("File not exist")
	}
}
