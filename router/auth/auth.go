package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	c.Set("token", token)
}
