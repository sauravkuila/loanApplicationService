package dao

type LoanDao interface {
	GetLoan() error
}

func (d *DAO) GetLoan() error {
	return nil
}
