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
