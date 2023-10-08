package killgrave

import (
	"bytes"
	"encoding/json"
	"github.com/pborman/uuid"
	"net/http"
	"time"
)

type BankAPIClient struct {
	URL        string
	HttpClient *http.Client
}

func NewKillGraveBankAPIClient(url string, timeout time.Duration) BankAPIClient {
	return BankAPIClient{
		URL: url,
		HttpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (cl *BankAPIClient) CreatePayment(req Payment) (Payment, error) {
	req.IdempotencyKey = uuid.New()

	body, err := json.Marshal(req)
	if err != nil {
		return Payment{}, err
	}

	request, err := http.NewRequest("POST", cl.URL+"/payments", bytes.NewBuffer(body))
	if err != nil {
		return Payment{}, err
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := cl.HttpClient.Do(request)
	if err != nil || response.StatusCode != 200 {
		return Payment{}, ErrBankProviderError
	}

	return req, nil
}

func (cl *BankAPIClient) GetPaymentByKey(key string) (Payment, error) {
	request, err := http.NewRequest("GET", cl.URL+"/payments", nil)
	if err != nil {
		return Payment{}, err
	}

	q := request.URL.Query()
	q.Add("key", key)
	request.URL.RawQuery = q.Encode()

	response, err := cl.HttpClient.Do(request)
	if err != nil || response.StatusCode != 200 {
		return Payment{}, ErrBankProviderError
	}

	var payments Payments
	err = json.NewDecoder(response.Body).Decode(&payments)
	if err != nil {
		return Payment{}, err
	}

	return payments.Payments[0], nil
}
