package test

import (
	"testing"

	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/external"
	"github.com/stretchr/testify/assert"
)

func TestCCExternalProvider_Positive(t *testing.T) {
	paymentGateway = PaymentGatewaySuccess{}
	assert.Equal(t, external.SUCCESS, paymentGateway.Validate(users[0].CCNo, users[0].CCExpDate))
}

func TestCCExternalProvider_Negative(t *testing.T) {
	paymentGateway = PaymentGatewayFail{}
	assert.Equal(t, external.FAIL, paymentGateway.Validate(users[0].CCNo, users[0].CCExpDate))
}

func TestChartServiceNotify_Positive(t *testing.T) {
	item := wareHouse.SearchByTitle("Go Ahead")
	if len(item) == 1 {
		qty := 1
		chartSvc = ChartServiceMockPositive{}
		notified := chartSvc.Notify(item[0].Title, item[0].Artist, qty)
		assert.Equal(t, true, notified)
	}
}

func TestChartServiceNotify_Negative(t *testing.T) {
	item := wareHouse.SearchByTitle("Go Ahead")
	if len(item) == 1 {
		qty := 1
		chartSvc = ChartServiceMockNegative{}
		notified := chartSvc.Notify(item[0].Title, item[0].Artist, qty)
		assert.Equal(t, false, notified)
	}
}