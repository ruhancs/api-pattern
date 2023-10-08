package entity

import (
	"errors"
	"time"

	"github.com/ruhancs/api-pattern/pkg/entity"
)

var (
	errIDRequired = errors.New("id is required")
	errIDInvalid = errors.New("id is invalid")
	errNameRequired = errors.New("name is required")
	errNameInvalid = errors.New("name is invalid")
	errPriceRequired = errors.New("price is required")
	errPriceInvalid = errors.New("price is invalid")
)

type Product struct {
	ID entity.ID `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product,error) {
	product := &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil,err
	}
	return product,nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return errIDRequired
	}
	if _,err := entity.ParseID(p.ID.String()); err != nil {
		return errIDInvalid
	}

	if p.Name == "" {
		return errNameRequired
	}
	if p.Price == 0 {
		return errPriceRequired
	}
	if p.Price < 0 {
		return errPriceInvalid
	}
	return nil
}