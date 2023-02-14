package main

import (
	"net/http"

	"github.com/lanpaiva/api/configs"

	"github.com/lanpaiva/api/internal/entity"
	"github.com/lanpaiva/api/internal/infra/database"
	"github.com/lanpaiva/api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	producthand := handlers.NewProductHandler(productDB)
	http.HandleFunc("/products", producthand.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
