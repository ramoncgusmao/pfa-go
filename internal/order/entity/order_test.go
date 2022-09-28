package entity_test

import (
	"testing"

	"github.com/ramoncgusmao/pfa-go/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateAnewOrder_ThenShouldReceiveAndError(t *testing.T) {

	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid id")

}

func TestGivenAnEmptyPrice_WhenCreateAnewOrder_ThenShouldReceiveAndError(t *testing.T) {

	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid price")

}

func TestGivenAnEmptyTax_WhenCreateAnewOrder_ThenShouldReceiveAndError(t *testing.T) {

	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid tax")

}

func TestGivenAValidParams_WhenCreateAnewOrder_ThenShouldReceiveAndError(t *testing.T) {

	order, err := entity.NewOrder("123", 10, 2)
	assert.NoError(t, err)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Equal(t, "123", order.ID)
}

func TestGivenAValidParams_WhenCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {

	order, err := entity.NewOrder("123", 10, 2)
	assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 12.0, order.FinalPrice)
}
