package debt

type DebtI interface {
	// AddDebt add debt
	AddDebt(debt Debt) error
	// GetDebt
	GetDebt(ulid string) (Debt, error)
	// GetDebts
	GetDebts() ([]Debt, error)
	// UpdateDebt
	UpdateDebt(debt Debt) error
	// DeleteDebt 删除债务
	DeleteDebt(ulid string) error
}
