package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/Adrianzzziny/FoodOrder/backend/database"
	"github.com/Adrianzzziny/FoodOrder/backend/models"
)

func ObtenerProductos(w http.ResponseWriter, r *http.Request) {

	//query para obtener productos por categoria
	categoria := r.URL.Query().Get("categoria")

	var productos []models.Producto
	query := database.DB

	if categoria != "" {
		query = query.Where("LOWER(categoria) = ?", strings.ToLower(categoria))
	}

	// Usamos la variable global 'db' que se setea en SetDB
	if err := query.Find(&productos).Error; err != nil {
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

func ObtenerCategorias(w http.ResponseWriter, r *http.Request) {
	var categorias []string

	// SELECT DISTINCT categoria FROM productos
	if err := database.DB.Model(&models.Producto{}).Distinct().Pluck("categoria", &categorias).Error; err != nil {
		http.Error(w, "Error al obtener categorías", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categorias)
}
