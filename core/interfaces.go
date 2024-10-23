package core

import "github.com/guilhermewolke/creditas-challenge-loans-go/dto"

type Loan interface {
	SetCustomer(customer dto.Customer)
	Evaluate() dto.LoanType
	isElegible() bool
}
