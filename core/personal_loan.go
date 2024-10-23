package core

import "github.com/guilhermewolke/creditas-challenge-loans-go/dto"

type PersonalLoan struct {
	customer dto.Customer
}

func (l *PersonalLoan) Evaluate() dto.LoanType {
	if l.isElegible() {
		return dto.LoanType{Type: dto.TYPE_PERSONAL_NAME,
			Tax: dto.TYPE_PERSONAL_FEE}
	}
	return dto.LoanType{}
}

func (l *PersonalLoan) SetCustomer(customer dto.Customer) {
	l.customer = customer
}

func (l *PersonalLoan) isElegible() bool {
	return true
}
