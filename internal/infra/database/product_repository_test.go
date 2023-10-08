package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ruhancs/api-pattern/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T){
	db,err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product,err := entity.NewProduct("P1", 10)
	assert.Nil(t,err)
	
	productRepo := NewProductRepository(db)
	err = productRepo.Create(product)
	assert.Nil(t,err)
	
	var productFounded entity.Product
	err = db.First(&productFounded,"id=?", product.ID).Error
	assert.Nil(t,err)
	assert.Equal(t, productFounded.ID,product.ID)
	assert.Equal(t, productFounded.Name,product.Name)
	assert.Equal(t, productFounded.Price,product.Price)
}

func TestFindAllProduct(t *testing.T){
	db,err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productRepo := NewProductRepository(db)
	for i := 1; i < 21; i++ {
		product,err := entity.NewProduct(fmt.Sprintf("P%d",i), rand.Intn(100)+1)
		assert.Nil(t,err)
		db.Create(product)
	}
	products,err := productRepo.FindAll(1,10,"asc")
	assert.Nil(t,err)
	assert.Len(t,products,10)
	assert.Equal(t,products[0].Name,"P1")
	assert.Equal(t,products[9].Name,"P10")
	
	products,err = productRepo.FindAll(2,10,"asc")
	assert.Nil(t,err)
	assert.Len(t,products,10)
	assert.Equal(t,products[0].Name,"P11")
	assert.Equal(t,products[9].Name,"P20")
}

func TestFindProductByID(t *testing.T){
	db,err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productRepo := NewProductRepository(db)
	
	product,err := entity.NewProduct("P1",10)
	assert.Nil(t,err)
	db.Create(product)
	
	productFounded,err := productRepo.FindByID(product.ID.String())
	assert.Nil(t,err)
	assert.Equal(t,product.ID,productFounded.ID)
	assert.Equal(t,product.Name,productFounded.Name)
	assert.Equal(t,product.Price,productFounded.Price)
}

func TestUpdateProduct(t *testing.T){
	db,err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productRepo := NewProductRepository(db)
	
	product,err := entity.NewProduct("P1",10)
	assert.Nil(t,err)
	db.Create(product)
	product.Name = "P updated"
	
	err = productRepo.Update(product)
	assert.Nil(t,err)

	var productFounded entity.Product
	err = db.First(&productFounded,"id=?", product.ID).Error
	assert.Nil(t,err)
	assert.NotNil(t,productFounded.ID)
	assert.Equal(t,productFounded.Name,"P updated")
	assert.Equal(t,productFounded.Price,product.Price)
}

func TestDeleteProduct(t *testing.T){
	db,err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productRepo := NewProductRepository(db)
	
	product,err := entity.NewProduct("P1",10)
	assert.Nil(t,err)
	db.Create(product)
	
	err = productRepo.Delete(product.ID.String())
	assert.Nil(t,err)

	var productFounded entity.Product
	err = db.First(&productFounded,"id=?", product.ID).Error
	assert.Error(t,err)
}