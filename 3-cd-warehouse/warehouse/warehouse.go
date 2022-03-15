package warehouse

import "github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/cd"

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
		if result := w.SearchByTitle(cd.Title); result == NOT_FOUND {
			w.stock = append(w.stock, cd)
		} else if result == SUCCESS {
			w.stock[i].Quantity += cd.Quantity
		}
	}
	return SUCCESS
}

func (w *Warehouse) SearchByTitle(title string) string {
	for _, cd := range w.stock {
		if cd.Title == title {
			return SUCCESS
		}
	}
	return NOT_FOUND
}

func (w *Warehouse) SearchByArtist(artist string) string {
	for _, cd := range w.stock {
		if cd.Artist == artist {
			return SUCCESS
		}
	}
	return NOT_FOUND
}

func (w *Warehouse) Rate(title string, rating int, comment string) string {
	if result := w.SearchByTitle(title); result == NOT_FOUND {
		return result
	}
	for _, item := range w.stock {
		if item.Title == title {
			item.Ratings = append(item.Ratings, cd.Rating{Rating: rating, Comment: comment})
			return SUCCESS
		}
	}
	return NOT_FOUND
}

func (w *Warehouse) Buy(title string, ccNo string, ccExpDate string) string {
	if result := w.SearchByTitle(title); result == NOT_FOUND {
		return result
	}
	for _, cd := range w.stock {
		if cd.Title == title {
			if cd.Quantity > 0 {
				cd.Quantity--
				return SUCCESS
			}
			return NOT_ENOUGH_STOCK
		}
	}
	return NOT_FOUND
}
