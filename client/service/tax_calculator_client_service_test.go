package service

import (
	"context"
	"errors"
	request "taxCalculator/client/sample_requests"
	"taxCalculator/mocks"
	"taxCalculator/proto"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxSuccessfully(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	taxCalculatorClient := mocks.NewMockTaxCalculatorServiceClient(mockCtrl)
	taxCalculatorClientService := NewTaxCalculatorClientService(taxCalculatorClient)
	sampleRequest := request.CreateTaxCalculatorRequest()
	response := &proto.TaxCalculatorResponse{
		Result: 123.45,
	}
	taxCalculatorClient.EXPECT().CalculateTax(
		context.TODO(), sampleRequest).Return(response, nil)
	err := taxCalculatorClientService.CalculateTax(sampleRequest)
	assert.Nil(t, err)
}

func TestCalculateTaxReturnsErrorForRPCCallError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	taxCalculatorClient := mocks.NewMockTaxCalculatorServiceClient(mockCtrl)
	taxCalculatorClientService := NewTaxCalculatorClientService(taxCalculatorClient)
	sampleRequest := request.CreateTaxCalculatorRequest()

	taxCalculatorClient.EXPECT().CalculateTax(
		context.TODO(), sampleRequest).Return(
		&proto.TaxCalculatorResponse{},
		errors.New("unable to connect to client"))
	err := taxCalculatorClientService.CalculateTax(sampleRequest)
	assert.NotNil(t, err)
}
