package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Adrianzzziny/FoodOrder/backend/models"
)

func CarritoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var input struct {
			ID int `json:"id"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		producto, encontrado := models.BuscarProductoPorID(input.ID)
		if !encontrado {
			http.Error(w, "Producto no encontrado en el menú", http.StatusNotFound)
			return
		}

		models.AgregarAlCarrito(producto)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"mensaje":  "Producto agregado al carrito",
			"producto": producto,
		})

		//Traer datos del carrito
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.Carrito)

		//Vaciar el carrito
	case http.MethodDelete:
		models.VaciarCarrito()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"mensaje": "Carrito vaciado",
		})

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
