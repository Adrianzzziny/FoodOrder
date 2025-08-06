package routes

import (
	"net/http"

	"github.com/Adrianzzziny/FoodOrder/backend/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("¡Bienvenido a Mi Pedido Fast!"))
	})
	http.HandleFunc("/menu", handlers.MenuHandler)
	http.HandleFunc("/carrito", handlers.CarritoHandler)
	http.HandleFunc("/pedido", handlers.PedidoHandler)
	http.HandleFunc("/pedido/listar", handlers.VerPedidosHandler)
}
