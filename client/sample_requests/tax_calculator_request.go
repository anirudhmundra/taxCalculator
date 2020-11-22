package request

import "taxCalculator/proto"

func CreateTaxCalculatorRequest() *proto.TaxCalculatorRequest {
	return &proto.TaxCalculatorRequest{
		Tds:                    100,
		RentalAllowance:        129,
		IncomeFromOtherSources: 134,
		LoanAmount:             1000,
		TotalIncome:            100000}
}
