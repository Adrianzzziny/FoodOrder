package models

type Pedido struct {
	Productos []Producto `json:"productos"`
	Total     float64    `json:"total"`
	Fecha     string     `json:"fecha"`
	Hora      string     `json:"hora"`
}

var PedidosRealizados []Pedido

func AgregarPedido(p Pedido) {
	PedidosRealizados = append(PedidosRealizados, p)
}
