package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch_ByTitlePositive(t *testing.T) {
	found := wareHouse.SearchByTitle("Go")
	assert.Equal(t, 3, len(found))
}

func TestSearch_ByTitleNegative(t *testing.T) {
	found := wareHouse.SearchByTitle("this doesnt exist")
	assert.Equal(t, 0, len(found))
}

func TestSearch_ByArtistPositive(t *testing.T) {
	found := wareHouse.SearchByArtist("Alicia Keys")
	assert.Equal(t, 2, len(found))
}
func TestSearch_ByArtistNegative(t *testing.T) {
	found := wareHouse.SearchByArtist("Jay Z")
	assert.Equal(t, 0, len(found))
}
