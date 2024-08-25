package handler

import (
	"github.com/gin-gonic/gin"
	"library-api/application/usecase"
	"library-api/domain"
	"net/http"
)

type UserHandler struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserHandler(uuc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: uuc}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var userRequest *domain.User

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.UserUseCase.CreateUser(userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh UserHandler) Login(c *gin.Context) {
	var userLogin *domain.User
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.UserUseCase.Login(*userLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := uh.UserUseCase.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var userRequest domain.User

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := uh.UserUseCase.UpdateUser(id, userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := uh.UserUseCase.GetUserByID(id)
	if err != nil {
		if err.Error() == "usuário não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
