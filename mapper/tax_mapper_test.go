package mapper

import (
	"taxCalculator/model"
	"taxCalculator/proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapTaxCalculatorRequestToModel(t *testing.T) {
	request := &proto.TaxCalculatorRequest{
		Tds:                    123,
		RentalAllowance:        12.34,
		IncomeFromOtherSources: 145.3,
		LoanAmount:             345,
		TotalIncome:            12324,
	}
	expectedTaxRequest := model.TaxRequest{
		TDS:                    123,
		RentalAllowance:        12.34,
		IncomeFromOtherSources: 145.3,
		LoanAmount:             345,
		TotalIncome:            12324,
	}
	taxRequest := MapTaxCalculatorRequestToModel(request)
	assert.Equal(t, expectedTaxRequest, taxRequest)
}
