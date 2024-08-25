package repository

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"library-api/domain"
)

type LoanRepository interface {
	Create(loan *domain.Loan) error
	Get(id string) (*domain.Loan, error)
	Delete(id string) error
	UpdateLoan(loan *domain.Loan) (*domain.Loan, error)
	ExistingEntity(model interface{}, id string) (bool, error)
}

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{db: db}
}

func (r loanRepository) Create(loan *domain.Loan) error {
	logrus.WithFields(logrus.Fields{
		"userid": loan.UserId,
		"bookid": loan.BookId,
	}).Info("Create new loan")

	return r.db.Create(&loan).Error
}

func (r loanRepository) Get(id string) (*domain.Loan, error) {
	var loan *domain.Loan
	if err := r.db.Where("id = ?", id).First(&loan).Error; err != nil {
		return loan, err
	}
	return loan, nil
}

func (r loanRepository) UpdateLoan(loan *domain.Loan) (*domain.Loan, error) {
	if err := r.db.Save(&loan).Error; err != nil {
		return loan, err
	}
	return loan, nil
}

func (r loanRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.Loan{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *loanRepository) ExistingEntity(model interface{}, id string) (bool, error) {
	result := r.db.Where("id = ?", id).First(model)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
