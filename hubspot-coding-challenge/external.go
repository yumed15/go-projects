package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type APIClient struct {
	HttpClient http.Client
	URL1       string
	URL2       string
	ApiKey     string
}

func NewGatewayClient(url1 string, url2 string, apiKey string) APIClient {
	return APIClient{
		HttpClient: http.Client{Timeout: time.Duration(1) * time.Minute},
		URL1:       url1,
		URL2:       url2,
		ApiKey:     apiKey,
	}
}

func (cl *APIClient) getData() (Partners, error) {

	req, err := http.NewRequest("GET", cl.URL1, nil)
	if err != nil {
		return Partners{}, err
	}

	q := req.URL.Query()
	q.Add("userKey", cl.ApiKey)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")

	response, err := cl.HttpClient.Do(req)

	if err != nil || response.StatusCode != http.StatusOK {
		return Partners{}, ErrAPIProviderError
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Partners{}, err
	}

	var res Partners
	err = json.Unmarshal(responseData, &res)
	if err != nil {
		return Partners{}, err
	}

	return res, nil
}

func (cl *APIClient) sendData(data BestDates) error {

	marshalled, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("failed to marshall data: %s", err)
	}

	req, err := http.NewRequest("POST", cl.URL2, bytes.NewReader(marshalled))
	if err != nil {
		fmt.Printf("failed to build request: %s", err)
	}

	q := req.URL.Query()
	q.Add("userKey", cl.ApiKey)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")

	res, err := cl.HttpClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return ErrAPIProviderError
	}

	fmt.Printf("Response %d %s", res.StatusCode, res.Body)

	return nil
}
