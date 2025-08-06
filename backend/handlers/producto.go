package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Adrianzzziny/FoodOrder/backend/database"
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

func ObtenerProductosPorID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var producto models.Producto
	if err := database.DB.First(&producto, id).Error; err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(producto)
}
