package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"library-api/domain"
)

type BookRepository interface {
	Create(book *domain.Book) error
	GetByID(id string) (*domain.Book, error)
	DeleteByID(id string) error
	UpdateBook(book *domain.Book) (*domain.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r bookRepository) Create(book *domain.Book) error {
	logrus.WithFields(logrus.Fields{
		"title": book.Title,
	}).Info("Create new book")

	return r.db.Create(&book).Error
}

func (r bookRepository) GetByID(id string) (*domain.Book, error) {
	var book *domain.Book
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (r bookRepository) DeleteByID(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.Book{}).Error; err != nil {
		return err
	}
	return nil
}

func (r bookRepository) UpdateBook(book *domain.Book) (*domain.Book, error) {
	if err := r.db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}
