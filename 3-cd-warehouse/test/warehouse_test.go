package test

import (
	"testing"

	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/cd"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestAddCDs_Positive(t *testing.T) {
	addcds := wareHouse.AddCDs([]cd.CD{
		{Title: "No stock", Artist: "Alicia Keys", Quantity: 10, Ratings: []cd.Rating{{Rating: 3, Comment: "no stock album"}}},
	})
	assert.Equal(t, warehouse.SUCCESS, addcds)
}

func TestAddCDs_NonExistant(t *testing.T) {
	addcds := wareHouse.AddCDs([]cd.CD{
		{Title: "this doesnt exists", Artist: "Alicia Keys", Quantity: 10, Ratings: []cd.Rating{{Rating: 3, Comment: "unexisting album"}}},
	})
	assert.Equal(t, warehouse.SUCCESS, addcds)
}
