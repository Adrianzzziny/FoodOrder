package database

import (
	"fmt"
	"log"

	"github.com/Adrianzzziny/FoodOrder/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=food_order_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	fmt.Println("Conexión exitosa a PostgreSQL 🎉")
	DB = db

	//Migracion Automatica
	db.AutoMigrate(&models.Producto{}, &models.Pedido{})
}

func InsertarProductosDePrueba() {

	var count int64
	DB.Model(&models.Producto{}).Count(&count)

	if count > 0 {

		log.Println("Productos ya existen en la base de datos")
		return
	}

	productos := []models.Producto{
		{Nombre: "Hamburguesa Clásica", Descripcion: "Carne y queso", Precio: 18.50, Categoria: "Comida", ImagenURL: "https://ejemplo.com/img/hamburguesa.jpg"},
		{Nombre: "Pizza Pepperoni", Descripcion: "Pizza con pepperoni", Precio: 25.00, Categoria: "Comida", ImagenURL: "https://ejemplo.com/img/pizza.jpg"},
		{Nombre: "Inca Kola 500ml", Descripcion: "Bebida Gasificada", Precio: 4.50, Categoria: "Bebida", ImagenURL: "https://ejemplo.com/img/incakola.jpg"},
		{Nombre: "Papas Fritas", Descripcion: "Fritas con salsa", Precio: 6.00, Categoria: "Complemento", ImagenURL: "https://ejemplo.com/img/papas.jpg"},
	}

	for _, p := range productos {
		result := DB.Create(&p)
		if result.Error != nil {
			log.Println("Error insertando producto:", p.Nombre, "-", result.Error)
		} else {
			log.Println("Producto insertado:", p.Nombre)
		}
	}
}
