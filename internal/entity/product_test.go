package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product,err := NewProduct("p1", 10)
	
	assert.Nil(t,err)
	assert.NotNil(t,product)
	assert.Equal(t,product.Name, "p1")
	assert.Equal(t,product.Price, 10)
	assert.NotEmpty(t,product.ID)
	assert.NotEmpty(t,product.CreatedAt)
}

func TestProductWhenNameISRequerid(t *testing.T) {
	product,err := NewProduct("", 10)
	
	assert.NotNil(t,err)
	assert.Error(t,err, "name is required")
	assert.Nil(t,product)
}

func TestProductWhenPriceISRequerid(t *testing.T) {
	product,err := NewProduct("p1", 0)
	
	assert.NotNil(t,err)
	assert.Error(t,err, "price is required")
	assert.Nil(t,product)
}

func TestProductWhenPriceISInvalid(t *testing.T) {
	product,err := NewProduct("p1", -1)
	
	assert.NotNil(t,err)
	assert.Error(t,err, "price is invalid")
	assert.Nil(t,product)
}

func TestProductValidate(t *testing.T) {
	product,err := NewProduct("p1", 10)
	
	assert.Nil(t,err)
	assert.Nil(t,product.Validate())
}