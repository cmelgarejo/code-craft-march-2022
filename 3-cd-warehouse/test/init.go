package test

import (
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/cd"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/customer"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/external"
	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/warehouse"
)

var (
	cds = []cd.CD{
		{Title: "No stock", Artist: "Alicia Keys", Rating: 3, Quantity: 0, Comments: []string{}},
		{Title: "Go Ahead", Artist: "Alicia Keys", Rating: 4, Quantity: 1, Comments: []string{}},
		{Title: "Ready 2 Go", Artist: "Martin Solveig", Rating: 5, Quantity: 4, Comments: []string{}},
		{Title: "Go", Artist: "Delilah", Rating: 1, Quantity: 9, Comments: []string{}},
	}
	users = []customer.Customer{
		{FullName: "John Doe", CCNo: "1234-1234-1234-1234", CCExpDate: "2025-01-01"},
		{FullName: "Jane Doe", CCNo: "2345-2345-2345-2345", CCExpDate: "2024-01-01"},
		{FullName: "John Smith", CCNo: "3456-3456-3456-3456", CCExpDate: "2023-01-01"},
	}
	wareHouse = warehouse.NewWarehouse(cds)
)

var paymentGateway external.PaymentGateway
