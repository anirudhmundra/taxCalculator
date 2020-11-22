package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxRequestValidatesSuccessfully(t *testing.T) {
	request := TaxRequest{
		TDS:                    100,
		RentalAllowance:        129,
		IncomeFromOtherSources: 134,
		LoanAmount:             1000,
		TotalIncome:            100000}

	assert.Nil(t, request.Validate())
}
func TestTaxRequestValidateForRentalAllowanced(t *testing.T) {
	request := TaxRequest{
		TDS:         100,
		TotalIncome: 100000}

	assert.Equal(t, errors.New("Invalid rental allowance"), request.Validate())
}

func TestTaxRequestValidateForTotalIncome(t *testing.T) {
	request := TaxRequest{
		TDS:             100,
		RentalAllowance: 129}
	assert.Equal(t, errors.New("Invalid total income"), request.Validate())
}

func TestTaxRequestValidateForTDS(t *testing.T) {
	request := TaxRequest{
		RentalAllowance: 129,
		TotalIncome:     100000}

	assert.Equal(t, errors.New("Invalid tds"), request.Validate())
}
