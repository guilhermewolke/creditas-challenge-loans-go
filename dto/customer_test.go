package dto

import "testing"

func TestBuild(t *testing.T) {
	builder := CustomerBuilder{}

	expectedName := "Teste 1"
	expectedCPF := "112.358.132-13"
	expectedAge := 30
	expectedLocation := "SP"
	expectedIncome := 5000

	builder.WithName(expectedName)

	customer := builder.Build()

	if customer.Name != expectedName {
		t.Fatalf("O nome esperado era %s mas o nome atual é %s", expectedName, customer.Name)
	}

	if customer.CPF != "" {
		t.Fatalf("O CPF esperado era %s mas o CPF atual é %s", "", customer.CPF)
	}

	if customer.Age != 0 {
		t.Fatalf("A idade esperada era %d mas a idade atual é %d", 0, customer.Age)
	}

	if customer.Location != "" {
		t.Fatalf("A localidade esperada era %s mas a localidade atual é %s", "", customer.Location)
	}

	if customer.Income != 0 {
		t.Fatalf("O salário esperado era %d mas o salário atual é %d", 0, customer.Income)
	}

	builder.WithCPF(expectedCPF).WithAge(expectedAge).WithLocation(expectedLocation).WithIncome(expectedIncome)

	if customer.Name != expectedName {
		t.Fatalf("O nome esperado era %s mas o nome atual é %s", expectedName, customer.Name)
	}

	if customer.CPF != expectedCPF {
		t.Fatalf("O CPF esperado era %s mas o CPF atual é %s", expectedCPF, customer.CPF)
	}

	if customer.Age != expectedAge {
		t.Fatalf("A idade esperada era %d mas a idade atual é %d", expectedAge, customer.Age)
	}

	if customer.Location != expectedLocation {
		t.Fatalf("A localidade esperada era %s mas a localidade atual é %s", expectedLocation, customer.Location)
	}

	if customer.Income != expectedIncome {
		t.Fatalf("O salário esperado era %d mas o salário atual é %d", expectedIncome, customer.Income)
	}

}
