package main

import (
	"fmt"
	"net/http"

	"github.com/Adrianzzziny/FoodOrder/backend/database"

	"github.com/Adrianzzziny/FoodOrder/backend/routes"
)

func main() {

	database.ConectarDB()

	routes.SetupRoutes()

	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
