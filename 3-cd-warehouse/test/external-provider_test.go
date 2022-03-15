package test

import (
	"testing"

	"github.com/cmelgarejo/code-craft-march-2022/3-cd-warehouse/external"
	"github.com/stretchr/testify/assert"
)

func TestCCExternalProvider_Positive(t *testing.T) {
	paymentGateway = external.PaymentGatewaySuccess{}
	assert.Equal(t, external.SUCCESS, paymentGateway.Validate(users[0].CCNo, users[0].CCExpDate))
}

func TestCCExternalProvider_Negative(t *testing.T) {
	paymentGateway = external.PaymentGatewayFail{}
	assert.Equal(t, external.FAIL, paymentGateway.Validate(users[0].CCNo, users[0].CCExpDate))
}
