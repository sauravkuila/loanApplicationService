package main

import (
	"context"

	v1 "github.com/aspire/v1"
	"github.com/gin-gonic/gin"
)

func initRouter() {
	gin.SetMode(gin.ReleaseMode)
}

func getRouter(ctx context.Context) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	// apis added in router
	v1AuthRoutes := router.Group("auth")
	{
		v1AuthRoutes.GET("login", v1.Login)
	}

	v1LoanApplicantRoutes := router.Group("loan")
	{
		v1LoanApplicantRoutes.POST("apply", v1.ApplyLoan)
		v1LoanApplicantRoutes.GET("status", v1.GetLoanStatus)
		v1LoanApplicantRoutes.GET("transactions", v1.GetPaymentTransaction)
		v1LoanApplicantRoutes.POST("transact", v1.ProcessLoanPayment)
	}

	v1LoanApproverRoutes := router.Group("approver/loan")
	{
		v1LoanApproverRoutes.GET("applicaitons", v1.GetPendingLoans)
		v1LoanApproverRoutes.POST("update", v1.UpdateLoanStatus)
	}

	return router
}
