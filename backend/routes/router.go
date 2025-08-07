package routes

import (
	"net/http"

	"github.com/Adrianzzziny/FoodOrder/backend/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("¡Bienvenido a Mi Pedido Fast!"))
	}).Methods("GET")

	r.HandleFunc("/menu", handlers.MenuHandler).Methods("GET")
	r.HandleFunc("/carrito", handlers.CarritoHandler).Methods("GET")
	r.HandleFunc("/pedido", handlers.PedidoHandler).Methods("POST")
	r.HandleFunc("/pedido/listar", handlers.VerPedidosHandler).Methods("GET")

	// Obtener productos desde la DB
	r.HandleFunc("/productos", handlers.ObtenerProductos).Methods("GET")

	// Obtener producto por ID desde la DB
	r.HandleFunc("/productos/{id}", handlers.ObtenerProductosPorID).Methods("GET")

	// Obtener categorias desde la DB
	r.HandleFunc("/categorias", handlers.ObtenerCategorias).Methods("GET")
}
