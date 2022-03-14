package external

type PaymentGateway interface {
	Validate(ccNo string, expDate string) string
}

type PaymentGatewaySuccess struct{}
type PaymentGatewayFail struct{}

var (
	SUCCESS = "payment success"
	FAIL    = "failed to process payment"
)

func (e PaymentGatewaySuccess) Validate(ccNo string, expDate string) string {
	return SUCCESS
}

func (e PaymentGatewayFail) Validate(ccNo string, expDate string) string {
	return FAIL
}
