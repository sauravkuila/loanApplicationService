package v1

import (
	"net/http"

	"github.com/aspire/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login models.Credentials
	err := c.ShouldBindJSON(&login)
	if err == nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ApplyLoan(c *gin.Context) {
	var loanRequest models.LoanRequest
	err := c.ShouldBindJSON(&loanRequest)
	if err == nil {
		//add to loan application list for user
		//return a loan application number
		c.JSON(http.StatusOK, nil)
	} else {
		//return relevant error
		c.JSON(http.StatusBadRequest, err)
	}
}

func GetLoanStatus(c *gin.Context) {
	//fetch loan number from query parameters
	//check if loan is assigned to respective user
	//return loan details

	//if no loan found, return error with invalid loan number
}

func GetPaymentTransaction(c *gin.Context) {
	//fetch loan number from query parameters
	//check if loan is assigned to respective user
	//return loan payment details
}

func ProcessLoanPayment(c *gin.Context) {
	var transact models.Transact
	err := c.ShouldBindJSON(&transact)
	if err == nil {
		//check loan id exists
		//check loan id is for user
		//validate loan for emi and outstanding
		c.JSON(http.StatusOK, nil)
	} else {
		//return relevant error
		c.JSON(http.StatusBadRequest, err)
	}
}

func GetPendingLoans(c *gin.Context) {
	//fetch all loans with details of user in APPLIED status
}

func UpdateLoanStatus(c *gin.Context) {
	//approve or reject loan
	var updateRequest models.UpdateLoanStatus
	err := c.ShouldBindJSON(&updateRequest)
	if err == nil {
		//check if loan exists
		//check if loan can change status
		c.JSON(http.StatusOK, nil)
	} else {
		//return relevant error
		c.JSON(http.StatusBadRequest, err)
	}
}
