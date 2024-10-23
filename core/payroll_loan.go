package core

import "github.com/guilhermewolke/creditas-challenge-loans-go/dto"

type PayrollLoan struct {
	customer dto.Customer
}

func (l *PayrollLoan) Evaluate() dto.LoanType {

	if l.isElegible() {
		return dto.LoanType{Type: dto.TYPE_PAYROLL_NAME,
			Tax: dto.TYPE_PAYROLL_FEE}
	}

	return dto.LoanType{}
}

func (l *PayrollLoan) SetCustomer(customer dto.Customer) {
	l.customer = customer
}

// Verifica se o cliente informado é elegível para este tipo de empréstimo
//
// Retorna true para elegível ou false para não elegível
func (l *PayrollLoan) isElegible() bool {
	return l.customer.Income >= 5000
}
