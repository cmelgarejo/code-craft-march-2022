package test

import (
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/cd"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/customer"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/external"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/warehouse"
)

type PaymentGatewaySuccess struct{}

type PaymentGatewayFail struct{}

func (e PaymentGatewaySuccess) Validate(ccNo string, expDate string) string {
	return external.SUCCESS
}

func (e PaymentGatewayFail) Validate(ccNo string, expDate string) string {
	return external.FAIL
}

type ChartServiceMockPositive struct{}

func (c ChartServiceMockPositive) Notify(title, artist string, qty int) bool {
	return true
}
func (c ChartServiceMockPositive) GetAlbum(title, artist string) external.Album {
	return external.Album{
		Title:       "Go Ahead",
		Artist:      "Alicia Keys",
		LowestPrice: 8,
	}
}

type ChartServiceMockNegative struct{}

func (c ChartServiceMockNegative) Notify(title, artist string, qty int) bool {
	return false
}
func (c ChartServiceMockNegative) GetAlbum(title, artist string) external.Album {
	return external.Album{}
}

var (
	paymentGateway external.PaymentGateway
	chartSvc       external.AlbumChartSvc
	cds            = []cd.CD{
		cd.NewCD("Go Ahead", "Alicia Keys", 1, 10, []cd.Rating{{Rating: 3, Comment: "Good album"}}),
		cd.NewCD("No stock", "Alicia Keys", 0, 11, []cd.Rating{{Rating: 3, Comment: "no stock album"}}),
		cd.NewCD("Ready 2 Go", "Martin Solveig", 8, 4, []cd.Rating{{Rating: 3, Comment: "ready album"}}),
		cd.NewCD("Go", "Delilah", 9, 5, []cd.Rating{{Rating: 3, Comment: "go there album"}}),
	}
	users = []customer.Customer{
		{FullName: "John Doe", CCNo: "1234-1234-1234-1234", CCExpDate: "2025-01-01"},
		{FullName: "Jane Doe", CCNo: "2345-2345-2345-2345", CCExpDate: "2024-01-01"},
		{FullName: "John Smith", CCNo: "3456-3456-3456-3456", CCExpDate: "2023-01-01"},
	}
	wareHouse = warehouse.NewWarehouse(cds)
)
