package test

import (
	"testing"

	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestBuy_Positive(t *testing.T) {
	item := wareHouse.SearchByTitle("Go Ahead")
	if len(item) == 1 {
		buy := item[0].Buy(users[0].CCNo, users[0].CCExpDate)
		assert.Equal(t, warehouse.SUCCESS, buy)
	}
}

func TestBuy_Negative(t *testing.T) {
	item := wareHouse.SearchByTitle("No stock")
	if len(item) == 1 {
		buy := item[0].Buy(users[0].CCNo, users[0].CCExpDate)
		assert.Equal(t, warehouse.NOT_ENOUGH_STOCK, buy)
	}
}

func TestRating_Positive(t *testing.T) {
	item := wareHouse.SearchByTitle("No stock")
	if len(item) == 1 {
		rate := item[0].Rate(5, "this album was not her best")
		assert.Equal(t, true, rate)
	}
}
