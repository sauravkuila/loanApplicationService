package business

import (
	"github.com/aspire/dao"
	"github.com/aspire/utils/database"
)

var loanContract LoanContract = nil

func GetLoanProcessor() LoanContract {
	if loanContract == nil {
		loanContract = &LoanProcessor{
			LoanDao: &dao.DAO{DB: database.Get()},
		}
	}
	return loanContract
}
