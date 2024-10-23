package core

import (
	"reflect"
	"testing"

	"github.com/guilhermewolke/creditas-challenge-loans-go/dto"
)

func TestCollateralizedIsElegible(t *testing.T) {
	builder := dto.CustomerBuilder{}
	loan := CollateralizedLoan{}

	// Idade 30
	// Mora em SP
	// Salário de 3000
	// Deve retornar falso
	builder.WithAge(30).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(3000)
	customer1 := builder.Build()
	loan.SetCustomer(*customer1)

	if loan.isElegible() {
		t.Fatalf("O valor esperado era falso, mas foi verdadeiro")
	}

	// Idade inferior a 30
	// Mora em SP
	// Salário de 3000
	// Deve retornar verdadeiro
	builder.WithAge(29).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(3000)
	customer2 := builder.Build()
	loan.SetCustomer(*customer2)

	if !loan.isElegible() {
		t.Fatalf("O valor esperado era verdadeiro, mas foi falso")
	}

	// Idade inferior a 30
	// Mora em BH
	// Salário de 3000
	// Deve retornar falso
	builder.WithAge(29).WithCPF("1123581321").WithLocation("BH").WithName("Cliente teste").WithIncome(3000)
	customer3 := builder.Build()
	loan.SetCustomer(*customer3)

	if loan.isElegible() {
		t.Fatalf("O valor esperado era falso, mas foi verdadeiro")
	}

	// Idade irrelevante
	// Mora em BH
	// Salário de 4000
	// Deve retornar falso
	builder.WithAge(29).WithCPF("1123581321").WithLocation("BH").WithName("Cliente teste").WithIncome(4000)
	customer4 := builder.Build()
	loan.SetCustomer(*customer4)

	if loan.isElegible() {
		t.Fatalf("O valor esperado era falso, mas foi verdadeiro")
	}

	// Idade irrelevante
	// Mora em SP
	// Salário de 4000
	// Deve retornar verdadeiro
	builder.WithAge(29).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(4000)
	customer5 := builder.Build()
	loan.SetCustomer(*customer5)

	if !loan.isElegible() {
		t.Fatalf("O valor esperado era verdadeiro, mas foi false")
	}

	// Idade 30
	// Localização irrelevante
	// Salário de 5000
	// Deve retornar false
	builder.WithAge(30).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(5000)
	customer6 := builder.Build()
	loan.SetCustomer(*customer6)

	if loan.isElegible() {
		t.Fatalf("O valor esperado era false, mas foi verdadeiro")
	}

	// Idade abaixo de 30
	// Localização irrelevante
	// Salário de 5000
	// Deve retornar verdadeiro
	builder.WithAge(29).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(5000)
	customer7 := builder.Build()
	loan.SetCustomer(*customer7)

	if !loan.isElegible() {
		t.Fatalf("O valor esperado era verdadeiro, mas foi false")
	}

}

func TestCollateralizedEvaluate(t *testing.T) {
	builder := dto.CustomerBuilder{}
	loan := CollateralizedLoan{}

	// Idade 30
	// Mora em SP
	// Salário de 3000
	// Deve retornar falso
	builder.WithAge(30).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(3000)
	customer1 := builder.Build()
	loan.SetCustomer(*customer1)
	result := loan.Evaluate()
	expected := dto.LoanType{}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("O valor esperado era %#v, mas o retorno foi %#v", expected, result)
	}

	// Idade inferior a 30
	// Mora em SP
	// Salário de 3000
	// Deve retornar verdadeiro
	builder.WithAge(29).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(3000)
	customer2 := builder.Build()
	loan.SetCustomer(*customer2)
	result = loan.Evaluate()
	expected = dto.LoanType{Type: dto.TYPE_COLLATERALIZED_NAME,
		Tax: dto.TYPE_COLLATERALIZED_FEE}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("O valor esperado era %#v, mas o retorno foi %#v", expected, result)
	}

	// Idade inferior a 30
	// Mora em BH
	// Salário de 3000
	// Deve retornar falso
	builder.WithAge(29).WithCPF("1123581321").WithLocation("BH").WithName("Cliente teste").WithIncome(3000)
	customer3 := builder.Build()
	loan.SetCustomer(*customer3)
	result = loan.Evaluate()
	expected = dto.LoanType{}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("O valor esperado era %#v, mas o retorno foi %#v", expected, result)
	}

	// Idade irrelevante
	// Mora em BH
	// Salário de 4000
	// Deve retornar falso
	builder.WithAge(29).WithCPF("1123581321").WithLocation("BH").WithName("Cliente teste").WithIncome(4000)
	customer4 := builder.Build()
	loan.SetCustomer(*customer4)

	result = loan.Evaluate()
	expected = dto.LoanType{}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("O valor esperado era %#v, mas o retorno foi %#v", expected, result)
	}

	// Idade irrelevante
	// Mora em SP
	// Salário de 4000
	// Deve retornar verdadeiro
	builder.WithAge(29).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(4000)
	customer5 := builder.Build()
	loan.SetCustomer(*customer5)

	result = loan.Evaluate()
	expected = dto.LoanType{Type: dto.TYPE_COLLATERALIZED_NAME,
		Tax: dto.TYPE_COLLATERALIZED_FEE}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("O valor esperado era %#v, mas o retorno foi %#v", expected, result)
	}

	// Idade 30
	// Localização irrelevante
	// Salário de 5000
	// Deve retornar false
	builder.WithAge(30).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(5000)
	customer6 := builder.Build()
	loan.SetCustomer(*customer6)

	result = loan.Evaluate()
	expected = dto.LoanType{}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("O valor esperado era %#v, mas o retorno foi %#v", expected, result)
	}

	// Idade abaixo de 30
	// Localização irrelevante
	// Salário de 5000
	// Deve retornar verdadeiro
	builder.WithAge(29).WithCPF("1123581321").WithLocation("SP").WithName("Cliente teste").WithIncome(5000)
	customer7 := builder.Build()
	loan.SetCustomer(*customer7)

	result = loan.Evaluate()
	expected = dto.LoanType{Type: dto.TYPE_COLLATERALIZED_NAME,
		Tax: dto.TYPE_COLLATERALIZED_FEE}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("O valor esperado era %#v, mas o retorno foi %#v", expected, result)
	}
}
