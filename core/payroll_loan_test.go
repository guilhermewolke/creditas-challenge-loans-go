package core

import (
	"reflect"
	"testing"

	"github.com/guilhermewolke/creditas-challenge-loans-go/dto"
)

func TestPayrollIsElegible(t *testing.T) {
	builder := dto.CustomerBuilder{}
	loan := PayrollLoan{}

	// Idade irrelevante
	// Localidade irrelevante
	// Salário de 3000 (abaixo de 5000)
	// Deve retornar falso
	builder.WithAge(30).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(3000)
	customer1 := builder.Build()
	loan.SetCustomer(*customer1)

	if loan.isElegible() {
		t.Fatalf("O valor esperado era falso, mas foi verdadeiro")
	}

	// Idade irrelevante
	// Localidade irrelevante
	// Salário de 5000 (igual ou superior à 5000)
	// Deve retornar verdadeiro
	builder.WithAge(29).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(5000)
	customer2 := builder.Build()
	loan.SetCustomer(*customer2)

	if !loan.isElegible() {
		t.Fatalf("O valor esperado era verdadeiro, mas foi falso")
	}

}

func TestPayrollEvaluate(t *testing.T) {
	builder := dto.CustomerBuilder{}
	loan := PayrollLoan{}

	// Idade irrelevante
	// Localidade irrelevante
	// Salário de 3000
	// Deve retornar vazio
	builder.WithAge(30).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(3000)
	customer1 := builder.Build()
	loan.SetCustomer(*customer1)
	result := loan.Evaluate()
	expected := dto.LoanType{}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("O valor esperado era %#v, mas o retorno foi %#v", expected, result)
	}

	// Idade irrelevante
	// Localidade irrelevante
	// Salário de 5000
	// Deve retornar um objeto preenchido
	builder.WithAge(29).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(5000)
	customer2 := builder.Build()
	loan.SetCustomer(*customer2)
	result = loan.Evaluate()
	expected = dto.LoanType{Type: dto.TYPE_PAYROLL_NAME,
		Tax: dto.TYPE_PAYROLL_FEE}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("O valor esperado era %#v, mas o retorno foi %#v", expected, result)
	}

}
