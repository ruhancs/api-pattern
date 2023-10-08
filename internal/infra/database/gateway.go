package database

import "github.com/ruhancs/api-pattern/internal/entity"

type UserRepositoryInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User,error)
}

type ProductRepositoryInterface interface {
	Create(product *entity.Product) error
	FindAll(limit int, sort string) ([]*entity.Product,error)
	FindByID(id string) (*entity.Product,error)
	Update(product *entity.Product) error
	Delete(id string) error
}