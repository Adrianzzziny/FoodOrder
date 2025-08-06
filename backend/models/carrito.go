package models

/*
// Producto representa un ítem del menú
type Producto struct {
	ID       string  `json:"id"`
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
}
*/

// Carrito es una lista de productos agregados por el usuario
var Carrito []Producto

// AgregarAlCarrito añade un producto al carrito en memoria
func AgregarAlCarrito(p Producto) {
	Carrito = append(Carrito, p)
}

// VaciarCarrito elimina todos los productos del carrito
func VaciarCarrito() {
	Carrito = []Producto{}
}

/*
Ya especificada con el metodo GET en handler/carrito.go
// VerCarritoHandler responde con el contenido del carrito en formato JSON
func VerCarritoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Carrito) // Ya estás dentro de models, no necesitas `models.Carrito`
}

*/
