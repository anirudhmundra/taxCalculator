package service

import (
	"context"
	"taxCalculator/proto"

	"github.com/sirupsen/logrus"
)

type taxCalculatorClientService struct {
	client proto.TaxCalculatorServiceClient
}

type TaxCalculatorClientService interface {
	CalculateTax(request *proto.TaxCalculatorRequest) error
}

func NewTaxCalculatorClientService(client proto.TaxCalculatorServiceClient) TaxCalculatorClientService {
	return &taxCalculatorClientService{client: client}
}

func (taxCalculatorClientService *taxCalculatorClientService) CalculateTax(
	request *proto.TaxCalculatorRequest) error {
	clientResponse, clientErr := taxCalculatorClientService.client.CalculateTax(
		context.Background(), request)
	if clientErr != nil {
		logrus.Errorf("client error occurred: %v", clientErr)
		return clientErr
	}
	logrus.Infof("Your tax amount is %v", clientResponse.GetResult())
	return nil
}
