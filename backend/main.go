package main

import (
	"fmt"
	"net/http"

	"github.com/Adrianzzziny/FoodOrder/backend/database"
	"github.com/Adrianzzziny/FoodOrder/backend/handlers"
	"github.com/Adrianzzziny/FoodOrder/backend/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.ConectarDB()
	handlers.SetDB(database.DB) // Le pasamos la conexión a los handlers

	database.InsertarProductosDePrueba()

	r := mux.NewRouter()
	routes.SetupRoutes(r)

	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
