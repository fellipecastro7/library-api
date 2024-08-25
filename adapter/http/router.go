package http

import (
	"github.com/gin-gonic/gin"
	"library-api/adapter/http/router"
	"library-api/container"
)

func SetupRouter(cont *container.Container) *gin.Engine {
	r := gin.Default()

	router.SetupUserRouters(r, cont.UserHandler)

	return r
}
