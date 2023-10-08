package database

import (
	"github.com/ruhancs/api-pattern/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) Create(user *entity.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) FindByEmail(email string) (*entity.User,error) {
	var user entity.User
	if err := repo.DB.Where("email=?", email).First(&user).Error; err != nil {
		return nil,err
	}
	return &user,nil
}