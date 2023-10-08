package main

import (
	"gateway/gateway"
	"log"
	"net/http"
)

func handleRequests() {
	gatewayClient := gateway.NewGatewayClient()

	http.HandleFunc("/token", gateway.AuthenticationHandler)
	http.HandleFunc("/payments", gatewayClient.CreatePaymentHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
