package service

import (
	"context"
	"taxCalculator/model"
	"taxCalculator/proto"

	"taxCalculator/constants"
	"taxCalculator/mapper"

	"github.com/sirupsen/logrus"
)

type taxCalculatorService struct {
}

func NewTaxCalculatorService() proto.TaxCalculatorServiceServer {
	return &taxCalculatorService{}
}

func (service *taxCalculatorService) CalculateTax(ctx context.Context,
	request *proto.TaxCalculatorRequest) (
	*proto.TaxCalculatorResponse, error) {
	taxRequest, err := mapAndValidate(request)
	if err != nil {
		logrus.Errorf("validation error occurred : %v", err)
		return &proto.TaxCalculatorResponse{}, err
	}
	return &proto.TaxCalculatorResponse{Result: calculateTax(taxRequest)}, nil
}

func calculateTax(
	request model.TaxRequest) float32 {
	//dummy logic
	totalRent := request.RentalAllowance * float32(0.5)
	totalLoanAmount := request.LoanAmount * constants.Inflation
	totalAmount := request.TotalIncome - totalRent - request.TDS +
		request.IncomeFromOtherSources - totalLoanAmount
	return totalAmount * getTaxPercent(totalAmount)
}

func mapAndValidate(
	request *proto.TaxCalculatorRequest) (
	model.TaxRequest, error) {

	taxRequest := mapper.MapTaxCalculatorRequestToModel(request)
	if validateRenterErr := taxRequest.Validate(); validateRenterErr != nil {
		logrus.Errorf("Validation error occurred for renter : %v", validateRenterErr)
		return model.TaxRequest{},
			validateRenterErr
	}
	return taxRequest, nil
}

func getTaxPercent(amount float32) float32 {
	if amount < 2500 {
		return 0
	} else if amount >= 2500 && amount < 4000 {
		return 0.30
	}
	return 0.40
}
