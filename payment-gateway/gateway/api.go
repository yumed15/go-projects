package gateway

import "gateway/banks/killgrave"

type BankInterface interface {
	CreatePayment(req killgrave.Payment) (killgrave.Payment, error)
	GetPaymentByKey(key string) (killgrave.Payment, error)
}
