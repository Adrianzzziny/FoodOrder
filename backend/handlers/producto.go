package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Adrianzzziny/FoodOrder/backend/models"
)

func ObtenerProductos(w http.ResponseWriter, r *http.Request) {
	var productos []models.Producto

	// Usamos la variable global 'db' que se setea en SetDB
	if err := db.Find(&productos).Error; err != nil {
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productos)
}
