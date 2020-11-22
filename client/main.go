package main

import (
	request "taxCalculator/client/sample_requests"
	"taxCalculator/client/service"
	"taxCalculator/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	client := proto.NewTaxCalculatorServiceClient(conn)
	taxCalculatorClientService := service.NewTaxCalculatorClientService(client)

	if err := taxCalculatorClientService.CalculateTax(
		request.CreateTaxCalculatorRequest()); err != nil {
		logrus.Errorf("error occurred: %v", err)
	}
}
