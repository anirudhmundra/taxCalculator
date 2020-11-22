package service

import (
	"context"
	"taxCalculator/proto"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TaxCalculatorServiceTest struct {
	suite.Suite
	service proto.TaxCalculatorServiceServer
}

func TestTaxCalculatorServiceSuite(t *testing.T) {
	suite.Run(t, new(TaxCalculatorServiceTest))
}

func (suite *TaxCalculatorServiceTest) SetupSuite() {
	suite.service = NewTaxCalculatorService()
}

func (suite *TaxCalculatorServiceTest) TestCalculateTaxWith40Percent() {
	request := &proto.TaxCalculatorRequest{
		Tds:                    100,
		RentalAllowance:        129,
		IncomeFromOtherSources: 134,
		LoanAmount:             1000,
		TotalIncome:            100000,
	}
	response, err := suite.service.CalculateTax(context.TODO(), request)
	suite.Nil(err)
	suite.Equal(float32(39955.2), response.Result)
}

func (suite *TaxCalculatorServiceTest) TestCalculateTaxWith0Percent() {
	request := &proto.TaxCalculatorRequest{
		Tds:                    100,
		RentalAllowance:        129,
		IncomeFromOtherSources: 134,
		LoanAmount:             100,
		TotalIncome:            100,
	}
	response, err := suite.service.CalculateTax(context.TODO(), request)
	suite.Nil(err)
	suite.Equal(float32(0), response.Result)
}

func (suite *TaxCalculatorServiceTest) TestCalculateTaxWith30Percent() {
	request := &proto.TaxCalculatorRequest{
		Tds:                    100,
		RentalAllowance:        29,
		IncomeFromOtherSources: 1840,
		LoanAmount:             1000,
		TotalIncome:            1000,
	}
	response, err := suite.service.CalculateTax(context.TODO(), request)
	suite.Nil(err)
	suite.Equal(float32(793.2), response.Result)
}

func (suite *TaxCalculatorServiceTest) TestCalculateTaxWhenTotalIncomeIsMissing() {
	request := &proto.TaxCalculatorRequest{
		Tds:             123,
		RentalAllowance: 12,
	}
	response, err := suite.service.CalculateTax(context.TODO(), request)
	suite.Equal("Invalid total income", err.Error())
	suite.Zero(response.Result)
}
