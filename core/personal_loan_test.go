package core

import (
	"reflect"
	"testing"

	"github.com/guilhermewolke/creditas-challenge-loans-go/dto"
)

func TestPersonalEvaluate(t *testing.T) {
	expected := dto.LoanType{
		Type: dto.TYPE_PERSONAL_NAME,
		Tax:  dto.TYPE_PERSONAL_FEE}

	var personalLoan PersonalLoan = PersonalLoan{}
	loanType := personalLoan.Evaluate()

	if !reflect.DeepEqual(expected, loanType) {
		t.Fatalf("O objeto esperado do tipo de empresto era %#v mas o retorno foi %#v", expected, loanType)
	}
}
