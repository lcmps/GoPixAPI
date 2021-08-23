

# GoPixAPI

A Go webapp capable of generating custom Pix QR Codes made and serving them in a RESTful API

## Endpoints

The following API endpoints are avaiable:
### `POST localhost:8080/qr`
`/qr` returns a PNG encoded image when provided with a JSON such as
```
 curl --request POST \
  --url http://localhost:8080/qr \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Eva Lu",
    "amount": 0,
    "city": "Sao Paulo",
    "description": "Hi!",
    "transactionId": "00001",
    "pixKey": "tasty@test.com",
    "foregroundColor": "#000000",
    "backgroundColor": "#00bf9b"
}'
```
### `POST localhost:8080/paste`
`/paste` returns a `text/plain` Pix copy-paste code when provided with a JSON such as
```
 curl --request POST \
  --url http://localhost:8080/paste \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Eva Lu",
    "amount": 0,
    "city": "Sao Paulo",
    "description": "Hi!",
    "transactionId": "00001",
    "pixKey": "tasty@test.com"
}'
```
Sample response:
```
 HTTP/1.1 200 OK
 Content-Type: text/plain
 00020126430014BR.GOV.BCB.PIX0114tasty@test.com0203Hi!52040000530398654040.005802BR5906Eva Lu6009Sao Paulo624305050000150300017BR.GOV.BCB.BRCODE01051.0.06304192F
```

### `POST localhost/link`
`/link` returns a `application/json` when provided with a JSON such as

```
 curl --request POST \
  --url http://localhost:8080/link \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Eva Lu",
    "amount": 0,
    "city": "Sao Paulo",
    "description": "Hi!",
    "transactionId": "00001",
    "pixKey": "tasty@test.com",
    "foregroundColor": "#000000",
    "backgroundColor": "#00bf9b"
}'
```
Sample response:
```
 {
	 "path": "/stored/2021-08-23-T14_tasty@test.com.png"
 }
```

## Build & Usage
Build the binary by running
`./scripts/init.sh` then `./GoPix`, default port is :8080

