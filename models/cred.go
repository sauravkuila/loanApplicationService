package models

const (
	LoanApprover  = 0
	LoanApplicant = 1
)

type UserInfo struct {
	Password string
	Username string
	Type     int
	Token    string
}
