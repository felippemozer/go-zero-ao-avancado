package secao10interfaces

import (
	cominterfaces "golang_zero_ao_avancado/secao10-interfaces/com-interfaces"
	"golang_zero_ao_avancado/secao10-interfaces/exemplos"
	interfacevazia "golang_zero_ao_avancado/secao10-interfaces/interface-vazia"
	seminterfaces "golang_zero_ao_avancado/secao10-interfaces/sem-interfaces"
)

func Main() {
	seminterfaces.Main()
	cominterfaces.Main()
	exemplos.Main()
	interfacevazia.Main()
}
