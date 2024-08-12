package model

import "time"

type Compra struct {
	DataCompra time.Time
	Mercado    *Mercado
	Itens      []*Item
}
