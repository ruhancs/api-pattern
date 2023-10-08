package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ruhancs/api-pattern/internal/dto"
	"github.com/ruhancs/api-pattern/internal/entity"
	"github.com/ruhancs/api-pattern/internal/infra/database"
)

type ProductHandler struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewProductHandler(repository database.ProductRepositoryInterface) *ProductHandler {
	return &ProductHandler{
		ProductRepository: repository,
	}
}

func(handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.ProductCreateDto
	err := json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product,err := entity.NewProduct(productDto.Name,productDto.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.ProductRepository.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusCreated)
}