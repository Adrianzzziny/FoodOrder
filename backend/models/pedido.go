package models

import (
	"time"

	"gorm.io/gorm"
)

type Pedido struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey" json:"id"`
	Productos []Producto `gorm:"many2many:pedido_productos;" json:"productos"`
	Total     float64    `json:"total"`
	Fecha     string     `json:"fecha"`
	Hora      string     `json:"hora"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

var PedidosRealizados []Pedido

func AgregarPedido(p Pedido) {
	PedidosRealizados = append(PedidosRealizados, p)
}
