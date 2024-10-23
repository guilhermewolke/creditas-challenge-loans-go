package dto

type Customer struct {
	Name     string
	CPF      string
	Age      int
	Location string
	Income   int
}

type CustomerBuilder struct {
	Customer Customer
}

func (c *CustomerBuilder) WithName(name string) *CustomerBuilder {
	c.Customer.Name = name
	return c
}

func (c *CustomerBuilder) WithCPF(cpf string) *CustomerBuilder {
	c.Customer.CPF = cpf
	return c
}

func (c *CustomerBuilder) WithAge(age int) *CustomerBuilder {
	c.Customer.Age = age
	return c
}

func (c *CustomerBuilder) WithLocation(location string) *CustomerBuilder {
	c.Customer.Location = location
	return c
}

func (c *CustomerBuilder) WithIncome(income int) *CustomerBuilder {
	c.Customer.Income = income
	return c
}

// Devolve o objeto criado à partir dos métodos auxiliares anteriores
//
// Retorna um ponteiro do tipo Customer
func (c *CustomerBuilder) Build() *Customer {
	return &c.Customer
}
