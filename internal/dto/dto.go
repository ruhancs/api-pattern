package dto

type ProductCreateDto struct {
	Name string `json:"name"`
	Price int `json:"price"`
}