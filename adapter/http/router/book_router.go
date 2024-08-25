package router

import (
	"github.com/gin-gonic/gin"
	"library-api/adapter/http/handler"
)

func SetupBookRouters(router *gin.Engine, bookHandler *handler.BookHandler) {
	bookGroup := router.Group("/api/v1/books")
	{
		bookGroup.POST("", bookHandler.CreateBook)
		bookGroup.GET("/:id", bookHandler.GetBookByID)
		bookGroup.PUT("/:id", bookHandler.UpdateBook)
		bookGroup.DELETE("/:id", bookHandler.DeleteBook)
	}
}
