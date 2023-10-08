## How to run the solution

- Run the payment gateway on http://localhost:10000
```
$ cd payment-gateway-yumed15
$ go build main.go
$ go run main.go
```

- Run the mock bank server on http://localhost:3000. The solution uses [`killgrave`](https://github.com/friendsofgo/killgrave) to create the bank mock server
```
# Install killgrave via brew
$ brew install friendsofgo/tap/killgrave

# Run the mock server using the defined configuration
$ cd payment-gateway-yumed15/bankmock
$ killgrave
```

## How to manually test a simple flow
- Authenticate `/token` using `"ClientKey": "key", "ClientSecret": "secret"`
```
curl --location --request GET 'http://localhost:10000/token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ClientKey": "key",
    "ClientSecret": "secret"
}'
```
![Get Token](static/token.png)

- Use the above token to create a payment request `/payments`
```
curl --location --request POST 'http://localhost:10000/payments' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU2NTU4MzIsInVzZXJuYW1lIjoia2V5In0.VAppBaJ5sEymnsJ-ouHmORgxrC_-zNOCLpKUSlEsCCw' \
--header 'Content-Type: application/json' \
--data-raw '{
  "cardNumber": "2321312",
  "expiryData": "102022",
  "amount": 15,
  "currency": "GBP",
  "cvv": "0123"
}'
```
![Create Payment](static/create-payment.png)

- Query the payment by key `/payments`
```
curl --location --request GET 'http://localhost:10000/payments?key=01D8EMQ185CA8PRGE20DKZTGSR' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU2NTU4MzIsInVzZXJuYW1lIjoia2V5In0.VAppBaJ5sEymnsJ-ouHmORgxrC_-zNOCLpKUSlEsCCw'
```
![Get payment by id](static/get-payment-by-id.png)


## Assumptions made
- We made the assumption that the merchant's solution is idempotent, so any incoming transaction request will be treated as a new payment request.
- We assume the acquiring bank provider is in charge of verifying the customer's account details and checks for sufficient funds. If any of those are not met then we assume we won't receive a 200.

## Areas for improvement
- We could consider the possibility that the merchant's solution is not idempotent which means we need to double-check if some transactions could be possible duplicates (due to network failures, etc). We could do this: 
  - either by refusing any second transaction from the same account made in a close time period to the first one. 
  - or requesting that the merchant provides an idempotency key with every request and ignore transactions with already existing idempotency keys.
- The configuration keys for authenticating the client should be securely stored in a vault solution.
- Set up an automated API documentation (possible solution could be swagger)
- There is currently no rate limiting neither for the payment gateway nor for the acquiring bank api, this should be added.
- Logging and appropriate monitoring should be added. 
- Incoming transaction requests should be encrypted and persisted. This means that when a merchant makes a request we need to 1. persist the transaction 2. make the request transaction to the acquiring bank. To ensure idempotency we could have an event driven architecture based on the transaction status.

merchant request comes in -> transaction persisted in append only storage with status `CREATED` -> consumer picks up transactions with the `CREATED` status, sends them to acquiring bank and updates the transaction status to `PROCESSED` or `FAILED` based on the response received.

## Suggested cloud technologies 
- When it comes to the cloud provider any is acceptable for the solution, but let's suppose we go ahead and use aws. 
- The solution could be run as a k8s cluster deployed inside an EKS.
- For the persistence storage, a sql solution such as aws RDS for PostgreSQL could be used to store and encrypt the incoming requests
- For logging and monitoring we could have logs pushed to elastic search and grafana used as a monitoring and alerting solution.