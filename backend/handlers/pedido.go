package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Adrianzzziny/FoodOrder/backend/models"
)

func PedidoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if len(models.Carrito) == 0 {
		http.Error(w, "El carrito está vacío", http.StatusBadRequest)
		return
	}

	// Calcular el total del pedido
	total := 0.0
	for _, producto := range models.Carrito {
		total += producto.Precio
	}

	// Obtener fecha y hora actual
	ahora := time.Now()
	fecha := ahora.Format("2006-01-02") // formato: AAAA-MM-DD
	hora := ahora.Format("15:04:05")    // formato: HH:MM:SS

	// Crear el pedido
	pedido := models.Pedido{
		Productos: models.Carrito,
		Total:     total,
		Fecha:     fecha,
		Hora:      hora,
	}

	// Guardar el pedido en la lista de pedidos
	models.AgregarPedido(pedido)

	// Vaciar el carrito después de confirmar el pedido
	models.Carrito = []models.Producto{}

	// Responder al cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensaje": "Pedido realizado con éxito",
		"pedido":  pedido,
	})
}

func VerPedidosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.PedidosRealizados)
}
