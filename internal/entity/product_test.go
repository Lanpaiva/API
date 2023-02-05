package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("product 1", 5000)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "product 1", product.Name)
	assert.Equal(t, 5000, product.Price)

}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 5000)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("product 1", 0)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}
func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("product 1", -1)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}
