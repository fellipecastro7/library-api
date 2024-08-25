package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"library-api/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	DeleteByID(id string) error
	UpdateUser(user *domain.User) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) Create(user *domain.User) error {
	logrus.WithFields(logrus.Fields{
		"username": user.Username,
		"email":    user.Email,
	}).Info("Create new user")

	return r.db.Create(&user).Error
}

func (r userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

func (r userRepository) GetByID(id string) (*domain.User, error) {
	var user *domain.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r userRepository) DeleteByID(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (r userRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
