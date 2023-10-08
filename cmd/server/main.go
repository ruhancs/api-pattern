package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ruhancs/api-pattern/internal/entity"
	"github.com/ruhancs/api-pattern/internal/infra/database"
	"github.com/ruhancs/api-pattern/internal/infra/webserver/handler"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	//config := configs.LoadConfig(".")
	db,err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) 
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	
	productRepository := database.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepository)

	r := chi.NewRouter()
	r.Post("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000",r)
}

