package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ruhancs/api-pattern/internal/dto"
	"github.com/ruhancs/api-pattern/internal/entity"
	"github.com/ruhancs/api-pattern/internal/infra/database"
	entityPkg "github.com/ruhancs/api-pattern/pkg/entity"
)

type ProductHandler struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewProductHandler(repository database.ProductRepositoryInterface) *ProductHandler {
	return &ProductHandler{
		ProductRepository: repository,
	}
}

// CreateProduct godoc
// @Summary      Create product
// @Description  create product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request   body      dto.ProductCreateDto  true  "product request"
// @Success      201
// @Failure      400  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products [post]
// @Security     ApiKeyAuth
func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.ProductCreateDto
	err := json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := entity.NewProduct(productDto.Name, productDto.Price)
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


// FindAll godoc
// @Summary      Find All products
// @Description  Find All products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page   query      string  false  "page number"
// @Param        limit  query      string  false  "limit"
// @Success      200  {array}   entity.Product
// @Failure      400  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products [get]
// @Security     ApiKeyAuth
func (handler *ProductHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query().Get("page")
	l := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")
	page, err := strconv.Atoi(p)
	if err != nil {
		page = 0
	}
	limit, err := strconv.Atoi(l)
	if err != nil {
		limit = 5
	}

	products, err := handler.ProductRepository.FindAll(page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProduct godoc
// @Summary      Get products
// @Description  Get products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "product id" Format(uuid)
// @Success      200  {object}   entity.Product
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [get]
// @Security     ApiKeyAuth
func (handler *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := handler.ProductRepository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)

}

// UpdateProduct godoc
// @Summary      Update products
// @Description  Update products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "product id" Format(uuid)
// @Param        request   body      dto.ProductCreateDto  true  "product request"
// @Success      204  
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [put]
// @Security     ApiKeyAuth
func (handler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.ProductRepository.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary      Delete products
// @Description  Delete products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "product id" Format(uuid)
// @Success      204  
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [delete]
// @Security     ApiKeyAuth
func (handler *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := handler.ProductRepository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
