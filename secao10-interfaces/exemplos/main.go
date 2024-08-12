package exemplos

import (
	"errors"
	"fmt"
)

func Main() {
	exibeError(errors.New("Teste"))

	p := ProblemaDeNetwork{
		rede:     true,
		hardware: false,
	}
	exibeError(p)
}

func exibeError(err error) {
	fmt.Println(err.Error())
}

type ProblemaDeNetwork struct {
	rede     bool
	hardware bool
}

func (p ProblemaDeNetwork) Error() string {
	if p.rede {
		return "Problema de rede"
	} else if p.hardware {
		return "Problema de hardware"
	}
	return "Outro problema"
}
