package router

import (
	"WhoaServer/router/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	user.SetupUserRouter(r)
	return r
}
