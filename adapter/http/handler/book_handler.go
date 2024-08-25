package handler

import (
	"github.com/gin-gonic/gin"
	"library-api/application/usecase"
	"library-api/domain"
	"net/http"
)

type BookHandler struct {
	BookUseCase *usecase.BookUseCase
}

func NewBookHandler(buc *usecase.BookUseCase) *BookHandler {
	return &BookHandler{BookUseCase: buc}
}

func (bh *BookHandler) CreateBook(c *gin.Context) {
	var bookRequest *domain.Book

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := bh.BookUseCase.CreateBook(bookRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (bh *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if err := bh.BookUseCase.DeleteBook(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livro deletado com sucesso"})
}

func (bh *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var bookRequest domain.Book

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook, err := bh.BookUseCase.UpdateBook(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

func (bh *BookHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	book, err := bh.BookUseCase.GetBookByID(id)
	if err != nil {
		if err.Error() == "livro não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, book)
}
