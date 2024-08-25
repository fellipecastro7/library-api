package router

import (
	"github.com/gin-gonic/gin"
	"library-api/adapter/http/handler"
)

func SetupUserRouters(router *gin.Engine, userHandler *handler.UserHandler) {
	userGroup := router.Group("/api/v1/users")
	{
		userGroup.POST("", userHandler.CreateUser)
		userGroup.POST("/login", userHandler.Login)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}
}
