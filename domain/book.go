package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	Id              uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Title           string         `json:"title" validate:"required,min=1,max=255"`
	Author          string         `json:"author" validate:"required,min=1,max=255"`
	PublicationYear *int           `json:"publicationYear" validate:"gte=0"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

func (b *Book) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
