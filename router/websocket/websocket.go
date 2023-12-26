package websocket

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func websocketHandler(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("messageType:", messageType)
		fmt.Println("p:", string(p))
		err = ws.WriteMessage(messageType, p)
		if err != nil {
			break
		}
	}
	ws.Close()
}

func SetupWebsocketRouter(r *gin.Engine) {
	r.GET("/ws", websocketHandler)
}
