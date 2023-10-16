package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPIClient(t *testing.T) {
	testCases := []struct {
		name             string
		serverStatusCode int
		serverResp       string
		expErr           error
	}{
		{
			name:             "Successful get data",
			serverStatusCode: http.StatusOK,
			serverResp:       `{"partners": []}`,
		},
		{
			name:             "Error on getting data",
			serverStatusCode: http.StatusBadGateway,
			expErr:           ErrAPIProviderError,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.serverStatusCode)
				w.Write([]byte(test.serverResp))
			}))
			defer srv.Close()

			cl := NewGatewayClient(srv.URL, srv.URL, "")
			_, err := cl.getData()
			if test.expErr != nil {
				require.EqualError(t, err, test.expErr.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestSendAPIClient(t *testing.T) {
	testCases := []struct {
		name             string
		serverStatusCode int
		serverResp       string
		expErr           error
	}{
		{
			name:             "Successful send data",
			serverStatusCode: http.StatusOK,
		},
		{
			name:             "Error on sending data",
			serverStatusCode: http.StatusBadGateway,
			expErr:           ErrAPIProviderError,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.serverStatusCode)
				w.Write([]byte(test.serverResp))
			}))
			defer srv.Close()

			cl := NewGatewayClient(srv.URL, srv.URL, "")
			err := cl.sendData(BestDates{})
			if test.expErr != nil {
				require.EqualError(t, err, test.expErr.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
