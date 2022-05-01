package models

const (
	Daily      = 0
	Weekly     = 1
	Monthly    = 2
	Quarterly  = 3
	HalfYearly = 4
	Yearly     = 5
)

type Credentials struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type LoanRequest struct {
	Amount             float64 `json:"amount" binding:"required"`
	Tenure             int64   `json:"tenure" binding:"required"`
	RepaymentFrequency int     `json:"-"`
}

type LoanResponse struct {
	Id          int64   `json:"id"`
	Amount      float64 `json:"amount"`
	Tenure      int64   `json:"tenure"`
	Frequency   string  `json:"frequency"`
	Emi         float64 `json:"emi"`
	Outstanding float64 `json:"outstanding"`
	Status      string  `json:"status"`
}

type UpdateLoanStatus struct {
	Id     int64  `json:"loanId" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type Transact struct {
	Id     int64   `json:"id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}
