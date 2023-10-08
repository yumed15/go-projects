package gateway

import (
	"encoding/json"
	"fmt"
	"gateway/banks/killgrave"
	"net/http"
	"time"
)

type BankClient struct {
	killGraveAPI     killgrave.BankAPIClient
	killGravePayment killgrave.Payment
}

func NewGatewayClient() BankClient {
	return BankClient{
		killGraveAPI: killgrave.NewKillGraveBankAPIClient(acquiringBankURL, 10*time.Second),
	}
}

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u Credentials
	json.NewDecoder(r.Body).Decode(&u)

	if u.ClientKey == clientKey && u.ClientSecret == clientSecret {
		tokenString, err := createToken(u.ClientKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("no username found")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "invalid credentials")
	}
}

func (cl *BankClient) CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(w, r) {
		return
	}

	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&cl.killGravePayment)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "invalid json body")
		}

		response, err := cl.killGraveAPI.CreatePayment(cl.killGravePayment)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
		}
		res, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
		}
		w.Write(res)
	case http.MethodGet:
		response, err := cl.killGraveAPI.GetPaymentByKey(r.URL.Query().Get("key"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
		}
		res, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
		}
		w.Write(res)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid rest method")
	}

}
