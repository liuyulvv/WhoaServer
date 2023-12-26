package router

import (
	"WhoaServer/router/auth"
	"WhoaServer/router/user"
	"WhoaServer/router/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(auth.AuthHandler)
	user.SetupUserRouter(r)
	websocket.SetupWebsocketRouter(r)
	return r
}
