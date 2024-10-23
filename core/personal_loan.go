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

// Verifica se o cliente informado é elegível para este tipo de empréstimo
//
// Retorna true para elegível ou false para não elegível
func (l *PersonalLoan) isElegible() bool {
	// Como no enunciado todos os clientes são elegíveis para empréstimos pessoais, então este método sempre retorna true.
	// Caso as regras de elegibilidade mudem para empréstimos pessoais, a alteração a ser feita na verificação é somente aqui
	return true
}
