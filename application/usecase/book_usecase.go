package usecase

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"library-api/adapter/repository"
	"library-api/domain"
	"time"
)

type BookUseCase struct {
	bookRepo repository.BookRepository
}

func NewBookUseCase(bookRepo repository.BookRepository) *BookUseCase {
	return &BookUseCase{bookRepo: bookRepo}
}

func (bc *BookUseCase) CreateBook(bookRequest *domain.Book) (*domain.Book, error) {
	existingBook, err := bc.bookRepo.GetByID(bookRequest.Id.String())
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if existingBook.Id != uuid.Nil {
		return nil, errors.New("livro já existe na biblioteca")
	}

	book := &domain.Book{
		Id:              uuid.New(),
		Title:           bookRequest.Title,
		Author:          bookRequest.Author,
		PublicationYear: bookRequest.PublicationYear,
		CreatedAt:       time.Now(),
	}

	err = book.Validate()
	if err != nil {
		return nil, err
	}

	if err = bc.bookRepo.Create(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (bc *BookUseCase) DeleteBook(id string) error {
	book, err := bc.bookRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("livro não encontrado")
		}
		return err
	}

	err = bc.bookRepo.DeleteByID(book.Id.String())
	if err != nil {
		return errors.New("erro ao deletar o livro")
	}

	return nil
}

func (bc *BookUseCase) UpdateBook(id string, bookRequest domain.Book) (*domain.Book, error) {
	existingBook, err := bc.bookRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("livro não encontrado")
		}
		return nil, err
	}

	if bookRequest.Title != "" {
		existingBook.Title = bookRequest.Title
	}

	if bookRequest.Author != "" {
		existingBook.Author = bookRequest.Author
	}

	if bookRequest.PublicationYear != nil {
		existingBook.PublicationYear = bookRequest.PublicationYear
	}

	if err = existingBook.Validate(); err != nil {
		return nil, err
	}

	updatedBook, err := bc.bookRepo.UpdateBook(existingBook)
	if err != nil {
		return nil, errors.New("erro ao atualizar o livro")
	}

	return updatedBook, nil
}

func (bc *BookUseCase) GetBookByID(id string) (*domain.Book, error) {
	book, err := bc.bookRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("livro não encontrado")
		}
		return nil, err
	}
	return book, nil
}
