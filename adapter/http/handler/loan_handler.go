package handler

import (
	"github.com/gin-gonic/gin"
	"library-api/application/usecase"
	"library-api/domain"
	"net/http"
)

type LoanHandler struct {
	LoanUseCase *usecase.LoanUseCase
}

func NewLoanHandler(luc *usecase.LoanUseCase) *LoanHandler {
	return &LoanHandler{LoanUseCase: luc}
}

func (lh *LoanHandler) CreateLoan(c *gin.Context) {
	var loanRequest *domain.Loan

	if err := c.ShouldBindJSON(&loanRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan, err := lh.LoanUseCase.CreateLoan(loanRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func (lh *LoanHandler) DeleteLoan(c *gin.Context) {
	id := c.Param("id")

	if err := lh.LoanUseCase.DeleteLoan(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registro de empréstimo excluído com sucesso"})
}

func (lh *LoanHandler) UpdateLoan(c *gin.Context) {
	id := c.Param("id")
	var loanRequest domain.Loan

	if err := c.ShouldBindJSON(&loanRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedLoan, err := lh.LoanUseCase.UpdateLoan(id, loanRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedLoan)
}

func (lh *LoanHandler) GetLoan(c *gin.Context) {
	id := c.Param("id")

	loan, err := lh.LoanUseCase.GetLoan(id)
	if err != nil {
		if err.Error() == "empréstimo não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, loan)
}
