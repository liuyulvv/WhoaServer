package router

import (
	"WhoaServer/router/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	user.SetupUserRouter(r)
	return r
}