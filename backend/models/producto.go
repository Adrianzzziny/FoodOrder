package models

import (
	"gorm.io/gorm"
)

type Producto struct {
	gorm.Model
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Categoria   string  `json:"categoria"`
	ImagenURL   string  `json:"imagen_url"`
	PedidoID    uint    `json:"pedido_id"` //Clave foranea
}

// Simulamos el menú disponible
var Menu = []Producto{
	{Model: gorm.Model{ID: 1}, Nombre: "Hamburguesa", Descripcion: "Carne y queso", Precio: 8.50, Categoria: "Comida"},
	{Model: gorm.Model{ID: 1}, Nombre: "Papas Fritas", Descripcion: "Fritas con salsa", Precio: 3.00, Categoria: "Comida"},
	{Model: gorm.Model{ID: 1}, Nombre: "Refresco", Descripcion: "Coca-Cola", Precio: 2.00, Categoria: "Bebida"},
}

// BuscarProductoPorID busca un producto por ID en el menú
func BuscarProductoPorID(id int) (Producto, bool) {

	for _, p := range Menu {
		if p.ID == uint(id) {
			return p, true
		}
	}
	return Producto{}, false
}
