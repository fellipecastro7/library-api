package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Loan struct {
	Id         uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserId     uuid.UUID      `json:"userId" gorm:"type:uuid;not null"`
	BookId     uuid.UUID      `json:"bookId" gorm:"type:uuid;not null"`
	LoanDate   time.Time      `json:"loanDate" gorm:"default:current_timestamp"`
	ReturnedAt *time.Time     `json:"returnedAt,omitempty"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

func (l *Loan) Validate() error {
	validate := validator.New()
	return validate.Struct(l)
}
