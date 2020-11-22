package model

import "errors"

type TaxRequest struct {
	TDS                    float32
	RentalAllowance        float32
	IncomeFromOtherSources float32
	LoanAmount             float32
	TotalIncome            float32
}

func (request *TaxRequest) Validate() error {
	if request.TDS <= 0 {
		return errors.New("Invalid tds")
	} else if request.RentalAllowance <= 0 {
		return errors.New("Invalid rental allowance")
	} else if request.TotalIncome <= 0 {
		return errors.New("Invalid total income")
	}
	return nil
}
