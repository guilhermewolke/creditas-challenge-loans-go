package dto

const (
	TYPE_PERSONAL_NAME       = "Pessoal"
	TYPE_PERSONAL_FEE        = 4
	TYPE_COLLATERALIZED_NAME = "Com Garantia"
	TYPE_COLLATERALIZED_FEE  = 3
	TYPE_PAYROLL_NAME        = "Consignado"
	TYPE_PAYROLL_FEE         = 2
)

type LoanOutput struct {
	Customer string
	Loans    []LoanType
}

type LoanType struct {
	Type string
	Tax  int
}
