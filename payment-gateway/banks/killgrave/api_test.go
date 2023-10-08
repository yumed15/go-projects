package killgrave

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBankAPIClient_CreatePayment(t *testing.T) {
	testCases := []struct {
		name             string
		serverStatusCode int
		serverResp       string
		request          Payment
		expErr           error
		expRes           Payment
	}{
		{
			name:             "Successful killGravePayment request",
			serverStatusCode: http.StatusOK,
			request: Payment{
				CardNumber: "312423423423",
				ExpiryDate: "102022",
				Amount:     10.21,
				Currency:   "GBP",
				CVV:        "123",
			},
		},
		{
			name:             "Error on bank provider error",
			serverStatusCode: http.StatusBadGateway,
			expErr:           ErrBankProviderError,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.serverStatusCode)
				w.Write([]byte(test.serverResp))
			}))
			defer srv.Close()

			cl := NewKillGraveBankAPIClient(srv.URL, 10*time.Second)
			_, err := cl.CreatePayment(test.request)
			if test.expErr != nil {
				require.EqualError(t, err, test.expErr.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestBankAPIClient_GetPaymentByKey(t *testing.T) {
	testCases := []struct {
		name             string
		serverStatusCode int
		serverResp       string
		request          string
		expErr           error
		expRes           Payment
	}{
		{
			name:             "Successful get killGravePayment by key",
			serverStatusCode: http.StatusOK,
			serverResp: `{
			  "payments": [
				{
				  "idempotencyKey": "01D8EMQ185CA8PRGE20DKZTGSR",
				  "cardNumber": "32132132",
				  "expiryDate": "102022",
				  "amount": 10.32,
				  "currency": "GBP",
				  "cvv": "021"
				}
			  ]
			}`,
			request: "01D8EMQ185CA8PRGE20DKZTGSR",
			expRes: Payment{
				IdempotencyKey: "01D8EMQ185CA8PRGE20DKZTGSR",
				CardNumber:     "32132132",
				ExpiryDate:     "102022",
				Amount:         10.32,
				Currency:       "GBP",
				CVV:            "021",
			},
		},
		{
			name:             "Error on bank provider error",
			serverStatusCode: http.StatusBadGateway,
			expErr:           ErrBankProviderError,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.serverStatusCode)
				w.Write([]byte(test.serverResp))
			}))
			defer srv.Close()

			cl := NewKillGraveBankAPIClient(srv.URL, 10*time.Second)
			response, err := cl.GetPaymentByKey(test.request)
			if test.expErr != nil {
				require.EqualError(t, err, test.expErr.Error())
			} else {
				require.Nil(t, err)
			}
			require.Equal(t, test.expRes, response)
		})
	}
}
