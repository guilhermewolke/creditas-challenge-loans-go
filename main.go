package main

import (
	"fmt"

	"github.com/guilhermewolke/creditas-challenge-loans-go/core"
	"github.com/guilhermewolke/creditas-challenge-loans-go/dto"
)

func main() {
	var (
		name, cpf, location string
		age, income         int
	)

	builder := dto.CustomerBuilder{}
	loan := core.LoanEvaluator{}

	fmt.Println("Insira o nome do cliente: ")
	fmt.Scanln(&name)

	fmt.Println("Insira o CPF do cliente: ")
	fmt.Scanln(&cpf)

	fmt.Println("Insira a Localidade do cliente: ")
	fmt.Scanln(&location)

	fmt.Println("Insira a idade do cliente: ")
	fmt.Scanln(&age)

	fmt.Println("Insira o salário do cliente: ")
	fmt.Scanln(&income)

	builder.WithAge(age).WithCPF(cpf).WithLocation(location).WithName(name).WithIncome(income)
	customer1 := builder.Build()
	result := loan.Execute(*customer1)

	fmt.Println("Resultado da análise:")
	fmt.Printf("%#v", result)

}
