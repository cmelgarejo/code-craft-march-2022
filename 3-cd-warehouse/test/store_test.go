package test

import (
	"testing"

	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/external"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/store"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestBuy_Positive(t *testing.T) {
	cd := cds[0]
	chartSvc = external.AlbumChartSvcImpl{}
	store := store.NewStore(wareHouse, chartSvc, paymentGateway)
	// album := store..GetAlbum(cd.Title, cd.Artist)
	result := store.Buy(cd, users[0].CCNo, users[0].CCExpDate, 1)
	assert.Equal(t, true, result.Notified)
	assert.Equal(t, warehouse.SUCCESS, result.Result)
	// assert.Less(t, sellingPrice, album.LowestPrice)
}

// func TestBuy_Negative(t *testing.T) {
// 	item := wareHouse.SearchByTitle("No stock")
// 	if len(item) == 1 {
// 		qty := 2
// 		result, _, _ := item[0].Buy(users[0].CCNo, users[0].CCExpDate, qty)
// 		assert.Equal(t, warehouse.NOT_ENOUGH_STOCK, result)
// 	}
// }
