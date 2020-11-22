package mapper

import (
	"taxCalculator/model"
	"taxCalculator/proto"
)

func MapTaxCalculatorRequestToModel(request *proto.TaxCalculatorRequest) model.TaxRequest {
	return model.TaxRequest{
		TDS:                    request.GetTds(),
		RentalAllowance:        request.GetRentalAllowance(),
		IncomeFromOtherSources: request.GetIncomeFromOtherSources(),
		LoanAmount:             request.GetLoanAmount(),
		TotalIncome:            request.GetTotalIncome(),
	}
}
