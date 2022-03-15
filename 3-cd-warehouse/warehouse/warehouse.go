package warehouse

import (
	"strings"

	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/cd"
)

const (
	NOT_ENOUGH_STOCK = "not enough stock"
	NOT_FOUND        = "not found"
	SUCCESS          = "success"
)

type Warehouse struct {
	stock []cd.CD
}

func NewWarehouse(cds []cd.CD) *Warehouse {
	return &Warehouse{
		stock: cds,
	}
}

func (w *Warehouse) AddCDs(cds []cd.CD) string {
	for i, cd := range cds {
		if result := w.SearchByTitle(cd.Title); len(result) == 0 {
			w.stock = append(w.stock, cd)
		} else if len(result) > 0 {
			w.stock[i].Quantity += cd.Quantity
		}
	}
	return SUCCESS
}

func (w *Warehouse) SearchByTitle(title string) []cd.CD {
	results := []cd.CD{}
	for _, cd := range w.stock {
		if strings.Contains(cd.Title, title) {
			results = append(results, cd)
		}
	}
	return results
}

func (w *Warehouse) SearchByArtist(artist string) []cd.CD {
	results := []cd.CD{}
	for _, cd := range w.stock {
		if strings.Contains(cd.Artist, artist) {
			results = append(results, cd)
		}
	}
	return results
}
