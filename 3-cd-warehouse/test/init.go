package test

import (
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/cd"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/customer"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/external"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/warehouse"
)

var (
	cds = []cd.CD{
		{Title: "No stock", Artist: "Alicia Keys", Quantity: 0, Ratings: []cd.Rating{{Rating: 3, Comment: "no stock album"}}},
		{Title: "Go Ahead", Artist: "Alicia Keys", Quantity: 1, Ratings: []cd.Rating{{Rating: 3, Comment: "Good album"}}},
		{Title: "Ready 2 Go", Artist: "Martin Solveig", Quantity: 4, Ratings: []cd.Rating{{Rating: 3, Comment: "ready album"}}},
		{Title: "Go", Artist: "Delilah", Quantity: 9, Ratings: []cd.Rating{{Rating: 3, Comment: "go there album"}}},
	}
	users = []customer.Customer{
		{FullName: "John Doe", CCNo: "1234-1234-1234-1234", CCExpDate: "2025-01-01"},
		{FullName: "Jane Doe", CCNo: "2345-2345-2345-2345", CCExpDate: "2024-01-01"},
		{FullName: "John Smith", CCNo: "3456-3456-3456-3456", CCExpDate: "2023-01-01"},
	}
	wareHouse = warehouse.NewWarehouse(cds)
)

var paymentGateway external.PaymentGateway
