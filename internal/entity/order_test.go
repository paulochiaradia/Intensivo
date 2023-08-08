package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfGetsAnErrorIfIdIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestIfGetsAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "test"}
	assert.Error(t, order.Validate(), "price need to be greater than zero")
}

func TestIfGetsAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{
		ID:    "test",
		Price: 10,
	}
	assert.Error(t, order.Validate(), "invalid tax")
}

func TestFinalPrice(t *testing.T) {
	order := Order{
		ID:    "132",
		Price: 10,
		Tax:   1,
	}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "132", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	err := order.CalculateFinalPrice()
	if err != nil {
		return
	}
	assert.Equal(t, 11.0, order.FinalPrice)

}
