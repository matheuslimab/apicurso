package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {

	product, err := NewProduct("Product 1", 0)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	assert.NotEmpty(t, product.Name)
	assert.NotEmpty(t, product.Price)

	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10, product.Price)

	assert.Nil(t, product.Validate())

}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -1)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}