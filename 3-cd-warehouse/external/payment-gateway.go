package external

type PaymentGateway interface {
	Validate(ccNo string, expDate string) string
}

var (
	SUCCESS = "payment success"
	FAIL    = "failed to process payment"
)