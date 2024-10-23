package core

import (
	"reflect"
	"testing"

	"github.com/guilhermewolke/creditas-challenge-loans-go/dto"
)

func TestExecute(t *testing.T) {
	//Cenário 1: O funcionário recebe até 3000, somente empréstimo pessoal é elegível para o cliente
	builder := dto.CustomerBuilder{}
	loan := LoanEvaluator{}

	// Idade 30
	// Mora em SP
	// Salário de 3000
	builder.WithAge(30).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(3000)
	customer1 := builder.Build()
	result := loan.Execute(*customer1)

	expected := dto.LoanOutput{
		Customer: "Cliente teste",
		Loans:    []dto.LoanType{{Type: dto.TYPE_PERSONAL_NAME, Tax: dto.TYPE_PERSONAL_FEE}}}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("O valor esperado era %#v mas o valor retornado foi %#v", expected, result)
	}

	//Cenário 2: O funcionário recebe até 3000, mas tem 30 anos e mora em SP, empréstimo pessoal e com garantia
	builder.WithAge(29).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(3000)
	customer2 := builder.Build()
	result = loan.Execute(*customer2)

	expected = dto.LoanOutput{
		Customer: "Cliente teste",
		Loans: []dto.LoanType{
			{Type: dto.TYPE_PERSONAL_NAME, Tax: dto.TYPE_PERSONAL_FEE},
			{Type: dto.TYPE_COLLATERALIZED_NAME, Tax: dto.TYPE_COLLATERALIZED_FEE}}}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("O valor esperado era %#v mas o valor retornado foi %#v", expected, result)
	}

	//Cenário 3: O funcionário recebe entre 3000 e 5000 de salário e não mora em SP, somente empréstimo pessoal é elegível para o cliente
	builder.WithAge(30).WithCPF("1123581321").WithLocation("BH").WithName("Cliente teste").WithIncome(4000)
	customer3 := builder.Build()
	result = loan.Execute(*customer3)
	expected = dto.LoanOutput{
		Customer: "Cliente teste",
		Loans:    []dto.LoanType{{Type: dto.TYPE_PERSONAL_NAME, Tax: dto.TYPE_PERSONAL_FEE}}}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("O valor esperado era %#v mas o valor retornado foi %#v", expected, result)
	}

	//Cenário 4: O funcionário recebe entre 3000 e 5000 de salário e mora em SP, empréstimo pessoal e com garantia são elegíveis para ele
	builder.WithAge(30).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(4000)
	customer4 := builder.Build()
	result = loan.Execute(*customer4)
	expected = dto.LoanOutput{
		Customer: "Cliente teste",
		Loans: []dto.LoanType{{Type: dto.TYPE_PERSONAL_NAME, Tax: dto.TYPE_PERSONAL_FEE},
			{Type: dto.TYPE_COLLATERALIZED_NAME, Tax: dto.TYPE_COLLATERALIZED_FEE}}}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("O valor esperado era %#v mas o valor retornado foi %#v", expected, result)
	}

	//Cenário 5: O funcionário recebe 5000 ou mais de salário, mas tem 30 anos ou mais, empréstimo pessoal e consignado são elegíveis para ele
	builder.WithAge(30).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(5000)
	customer5 := builder.Build()
	result = loan.Execute(*customer5)
	expected = dto.LoanOutput{
		Customer: "Cliente teste",
		Loans: []dto.LoanType{{Type: dto.TYPE_PERSONAL_NAME, Tax: dto.TYPE_PERSONAL_FEE},
			{Type: dto.TYPE_PAYROLL_NAME, Tax: dto.TYPE_PAYROLL_FEE}}}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("O valor esperado era %#v mas o valor retornado foi %#v", expected, result)
	}

	//Cenário 6: O funcionário recebe 5000 ou mais de salário, e tem menos de 30 anos, todas as modalidades de empréstimo são elegíveis para ele
	builder.WithAge(29).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(5000)
	customer6 := builder.Build()
	result = loan.Execute(*customer6)
	expected = dto.LoanOutput{
		Customer: "Cliente teste",
		Loans: []dto.LoanType{{Type: dto.TYPE_PERSONAL_NAME, Tax: dto.TYPE_PERSONAL_FEE},
			{Type: dto.TYPE_COLLATERALIZED_NAME, Tax: dto.TYPE_COLLATERALIZED_FEE},
			{Type: dto.TYPE_PAYROLL_NAME, Tax: dto.TYPE_PAYROLL_FEE}}}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("O valor esperado era %#v mas o valor retornado foi %#v", expected, result)
	}
}
