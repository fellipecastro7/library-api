package router

import (
	"github.com/gin-gonic/gin"
	"library-api/adapter/http/handler"
)

func SetupLoanRouters(router *gin.Engine, loanHandler *handler.LoanHandler) {
	loanGroup := router.Group("/api/v1/loans")
	{
		loanGroup.POST("", loanHandler.CreateLoan)
		loanGroup.GET("/:id", loanHandler.GetLoan)
		loanGroup.PUT("/:id", loanHandler.UpdateLoan)
		loanGroup.DELETE("/:id", loanHandler.DeleteLoan)
	}
}
