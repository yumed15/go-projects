package killgrave

type Payment struct {
	IdempotencyKey string  `json:"idempotencyKey"`
	CardNumber     string  `json:"cardNumber"`
	ExpiryDate     string  `json:"expiryDate"` // monthyear (e.g. 102022)
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	CVV            string  `json:"cvv"`
}

type Payments struct {
	Payments []Payment `json:"payments"`
}
