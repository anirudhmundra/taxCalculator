grpc-server:
	go run cmd/server/main.go -port 8080

rest-server:
	go run cmd/rest/main.go -port 8081

grpc-client:
	go run client/main.go

test:
	go test -cover -race ./...

build:
	go build ./...

generate-proto:
	cd proto/
	protoc --go_out=plugins=grpc:. taxCalculatorService.proto
	protoc --go_out=plugins=grpc:. taxCalculatorService.proto --grpc-gateway_out=:.