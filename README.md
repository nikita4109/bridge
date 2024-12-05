## Usage

To run the service, use the following command:

`
PKEY=$YOUR_PRIVATE_KEY docker compose up -d --build
`

For testing, I have implemented a client. To run the client, use:

`
go run cmd/client/main.go
`
