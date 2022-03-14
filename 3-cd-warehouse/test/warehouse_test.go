package test

import (
	"testing"

	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/cd"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/external"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestBuy_Positive(t *testing.T) {
	buy := wareHouse.Buy("Go Ahead", users[0].CCNo, users[0].CCExpDate)
	assert.Equal(t, warehouse.SUCCESS, buy)
}

func TestBuy_Negative(t *testing.T) {
	buy := wareHouse.Buy("No stock", users[1].CCNo, users[1].CCExpDate)
	assert.Equal(t, warehouse.NOT_ENOUGH_STOCK, buy)
}

func TestCCExternalProvider_Positive(t *testing.T) {
	paymentGateway = external.PaymentGatewaySuccess{}
	assert.Equal(t, external.SUCCESS, paymentGateway.Validate(users[0].CCNo, users[0].CCExpDate))
}

func TestCCExternalProvider_Negative(t *testing.T) {
	paymentGateway = external.PaymentGatewayFail{}
	assert.Equal(t, external.FAIL, paymentGateway.Validate(users[0].CCNo, users[0].CCExpDate))
}

func TestAddCDs_Positive(t *testing.T) {
	rate := wareHouse.AddCDs([]cd.CD{
		{Title: "No stock", Artist: "Alicia Keys", Rating: 3, Quantity: 10, Comments: []string{}},
	})
	assert.Equal(t, warehouse.SUCCESS, rate)
}

func TestAddCDs_Negative(t *testing.T) {
	rate := wareHouse.AddCDs([]cd.CD{
		{Title: "this doesnt exists", Artist: "Alicia Keys", Rating: 3, Quantity: 10, Comments: []string{}},
	})
	assert.Equal(t, warehouse.SUCCESS, rate)
}
