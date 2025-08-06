package models

type Producto struct {
	ID        int     `json:"id"`
	Nombre    string  `json:"nombre"`
	Precio    float64 `json:"precio"`
	Categoria string  `json:"categoria"`
}

// Simulamos el menú disponible
var Menu = []Producto{
	{ID: 1, Nombre: "Hamburguesa", Precio: 8.50, Categoria: "Comida"},
	{ID: 2, Nombre: "Papas Fritas", Precio: 3.00, Categoria: "Comida"},
	{ID: 3, Nombre: "Refresco", Precio: 2.00, Categoria: "Bebida"},
}

// BuscarProductoPorID busca un producto por ID en el menú
func BuscarProductoPorID(id int) (Producto, bool) {

	for _, p := range Menu {
		if p.ID == id {
			return p, true
		}
	}
	return Producto{}, false
}
