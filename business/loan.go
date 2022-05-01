package business

import (
	"github.com/aspire/dao"
	"github.com/aspire/models"
)

type LoanProcessor struct {
	LoanDao dao.LoanDao
}

type LoanContract interface {
	CheckAndApply(request models.LoanRequest) error
}

func (p *LoanProcessor) CheckAndApply(request models.LoanRequest) error {
	return nil
}
