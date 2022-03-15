package store

import (
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/cd"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/external"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/warehouse"
)

type Store struct {
	cdWarehouse    *warehouse.Warehouse
	chartSvc       external.AlbumChartSvc
	paymentGateway external.PaymentGateway
}

type BuyResult struct {
	Result   string
	SoldFor  int
	Notified bool
}

func NewStore(cdWarehouse *warehouse.Warehouse, chartSvc external.AlbumChartSvc, paymentGateway external.PaymentGateway) *Store {
	return &Store{
		cdWarehouse:    cdWarehouse,
		chartSvc:       chartSvc,
		paymentGateway: paymentGateway,
	}
}

func (store *Store) Buy(item cd.CD, ccNo string, ccExpDate string, qty int) BuyResult {
	store.cdWarehouse.AddCDs([]cd.CD{})
	if item.Quantity > 0 && item.Quantity-qty >= 0 {
		item.Quantity -= qty
		album := store.chartSvc.GetAlbum(item.Title, item.Artist)
		lowestPrice := item.Price
		if album.LowestPrice < item.Price {
			lowestPrice = album.LowestPrice - 1
		}
		notified := store.chartSvc.Notify(item.Title, item.Artist, qty)
		return BuyResult{
			Result:   warehouse.SUCCESS,
			SoldFor:  lowestPrice,
			Notified: notified,
		}
	}
	// return NOT_ENOUGH_STOCK, item.Price, false//
	return BuyResult{
		Result:   warehouse.NOT_ENOUGH_STOCK,
		SoldFor:  item.Price,
		Notified: false,
	}
}
