package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id           uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Username     string         `json:"username" validate:"required,min=1,max=255"`
	Email        string         `json:"email" validate:"required,email,max=255"`
	Password     string         `json:"password" validate:"required" gorm:"-"`
	PasswordHash string         `json:"-" validate:"required,min=1,max=255"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
