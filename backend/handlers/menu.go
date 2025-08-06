package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Adrianzzziny/FoodOrder/backend/models"
)

func MenuHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Menu)
}
