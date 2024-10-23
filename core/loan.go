package core

import (
	"github.com/guilhermewolke/creditas-challenge-loans-go/dto"
)

type LoanEvaluator struct {
	Customer dto.Customer
}

func (l *LoanEvaluator) Execute(customer dto.Customer) dto.LoanOutput {
	output := dto.LoanOutput{}
	output.Customer = customer.Name

	loanTypes := make([]dto.LoanType, 0, 3)

	// Tratando o empréstimo pessoal
	personalLoan := PersonalLoan{}
	personalLoan.SetCustomer(customer)

	loanType := personalLoan.Evaluate()

	if loanType.Type != "" {
		loanTypes = append(loanTypes, loanType)
	}

	// Tratando o empréstimo com garantia
	collateralizedLoan := CollateralizedLoan{}
	collateralizedLoan.SetCustomer(customer)

	loanType = collateralizedLoan.Evaluate()

	if loanType.Type != "" {
		loanTypes = append(loanTypes, loanType)
	}

	// Tratando o empréstimo consignado
	payrollLoan := PayrollLoan{}
	payrollLoan.SetCustomer(customer)

	loanType = payrollLoan.Evaluate()

	if loanType.Type != "" {
		loanTypes = append(loanTypes, loanType)
	}

	output.Loans = loanTypes

	return output
}
