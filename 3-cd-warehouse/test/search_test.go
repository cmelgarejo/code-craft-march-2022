package test

import (
	"testing"

	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestSearch_ByTitlePositive(t *testing.T) {
	found := wareHouse.SearchByTitle("Go")
	assert.Equal(t, warehouse.SUCCESS, found)
}

func TestSearch_ByTitleNegative(t *testing.T) {
	found := wareHouse.SearchByTitle("this doesnt exist")
	assert.Equal(t, warehouse.NOT_FOUND, found)
}

func TestSearch_ByArtistPositive(t *testing.T) {
	found := wareHouse.SearchByArtist("Alicia Keys")
	assert.Equal(t, warehouse.SUCCESS, found)
}
func TestSearch_ByArtistNegative(t *testing.T) {
	found := wareHouse.SearchByArtist("Jay Z")
	assert.Equal(t, warehouse.NOT_FOUND, found)
}
