package database

import (
	"github.com/ruhancs/api-pattern/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (repository *ProductRepository) Create(product *entity.Product) error {
	return repository.DB.Create(product).Error
}

func (repository *ProductRepository) FindAll(limit int, sort string) ([]*entity.Product,error) {
	return nil,nil
}

func (repository *ProductRepository) FindByID(id string) (*entity.Product,error) {
	var product entity.Product
	err := repository.DB.First(&product,"id=?",id).Error
	return &product,err
}

func (repository *ProductRepository) Update(product *entity.Product) error {
	_,err := repository.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return repository.DB.Save(product).Error
}

func (repository *ProductRepository) Delete(id string) error {
	product,err := repository.FindByID(id)
	if err != nil {
		return err
	}
	return repository.DB.Delete(product).Error
}

