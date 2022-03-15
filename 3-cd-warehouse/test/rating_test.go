package test

import (
	"testing"

	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestRating_Positive(t *testing.T) {
	rate := wareHouse.Rate("No stock", 5, "this album was not her best")
	assert.Equal(t, warehouse.SUCCESS, rate)
}

func TestRating_Negative(t *testing.T) {
	rate := wareHouse.Rate("this doesnt Exists", 5, "this album was not her best")
	assert.Equal(t, warehouse.NOT_FOUND, rate)
}
