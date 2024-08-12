package cominterfaces

type geometria interface {
	area() float64
}

type retangulo struct {
	largura, altura float64
}

type circulo struct {
	raio float64
}
