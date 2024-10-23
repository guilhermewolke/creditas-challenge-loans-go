package core

import "github.com/guilhermewolke/creditas-challenge-loans-go/dto"

type CollateralizedLoan struct {
	customer dto.Customer
}

func (l *CollateralizedLoan) Evaluate() dto.LoanType {

	if l.isElegible() {
		return dto.LoanType{Type: dto.TYPE_COLLATERALIZED_NAME,
			Tax: dto.TYPE_COLLATERALIZED_FEE}
	}

	return dto.LoanType{}
}

func (l *CollateralizedLoan) SetCustomer(customer dto.Customer) {
	l.customer = customer
}

func (l *CollateralizedLoan) isElegible() bool {
	switch {
	case l.customer.Income <= 3000:
		return l.customer.Age < 30 && l.customer.Location == "SP"
	case l.customer.Income > 3000 && l.customer.Income < 5000:
		return l.customer.Location == "SP"
	case l.customer.Income >= 5000:
		return l.customer.Age < 30
	default:
		return false
	}
}
