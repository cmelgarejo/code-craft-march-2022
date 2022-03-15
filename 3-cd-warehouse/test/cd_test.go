package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRating_Positive(t *testing.T) {
	item := wareHouse	.SearchByTitle("No stock")
	if len(item) == 1 {
		rate := item[0].Rate(5, "this album was not her best")
		assert.Equal(t, true, rate)
	}
}
