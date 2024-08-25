package usecase

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"library-api/adapter/repository"
	"library-api/domain"
	"time"
)

type LoanUseCase struct {
	loanRepo repository.LoanRepository
}

func NewLoanUseCase(loanRepo repository.LoanRepository) *LoanUseCase {
	return &LoanUseCase{loanRepo: loanRepo}
}

func (lc *LoanUseCase) CreateLoan(loanRequest *domain.Loan) (*domain.Loan, error) {
	//userExists, err := lc.userRepo.GetByID(loanRequest.UserId.String())
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		return nil, errors.New("usuário não encontrado")
	//	}
	//	return nil, err
	//}
	//
	//if userExists == nil {
	//	return nil, errors.New("usuário não encontrado")
	//}
	//
	//bookExists, err := lc.bookRepo.GetByID(loanRequest.BookId.String())
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		return nil, errors.New("livro não encontrado")
	//	}
	//	return nil, err
	//}
	//if bookExists == nil {
	//	return nil, errors.New("livro não encontrado")
	//}

	var (
		user *domain.User
		book *domain.Book
	)

	existUser, _ := lc.loanRepo.ExistingEntity(&user, loanRequest.UserId.String())
	if !existUser {
		return nil, errors.New("erro: usuario não encontrado")
	}

	existBook, _ := lc.loanRepo.ExistingEntity(&book, loanRequest.BookId.String())
	if !existBook {
		return nil, errors.New("erro: livro não encontrado")
	}

	loan := &domain.Loan{
		Id:       uuid.New(),
		UserId:   loanRequest.UserId,
		BookId:   loanRequest.BookId,
		LoanDate: time.Now(),
	}

	if err := loan.Validate(); err != nil {
		return nil, err
	}

	err := lc.loanRepo.Create(loan)
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (lc *LoanUseCase) GetLoan(id string) (*domain.Loan, error) {
	loan, err := lc.loanRepo.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("empréstimo não encontrado")
		}
		return nil, err
	}
	return loan, nil
}

func (lc *LoanUseCase) UpdateLoan(id string, loanRequest domain.Loan) (*domain.Loan, error) {
	existingLoan, err := lc.loanRepo.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("empréstimo não encontrado")
		}
		return nil, err
	}

	if loanRequest.ReturnedAt != nil {
		existingLoan.ReturnedAt = loanRequest.ReturnedAt
	}

	if err := existingLoan.Validate(); err != nil {
		return nil, err
	}

	updatedLoan, err := lc.loanRepo.UpdateLoan(existingLoan)
	if err != nil {
		return nil, errors.New("erro ao atualizar o empréstimo")
	}

	return updatedLoan, nil
}

func (lc *LoanUseCase) DeleteLoan(id string) error {
	loan, err := lc.loanRepo.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("empréstimo não encontrado")
		}
		return err
	}

	err = lc.loanRepo.Delete(loan.Id.String())
	if err != nil {
		return errors.New("erro ao deletar o empréstimo")
	}

	return nil
}
