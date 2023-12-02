package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func loginHandler(c *gin.Context) {

	secret := []byte("secret")

	var json struct {
		User     string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&json) == nil {
		if json.User == "liuyulvv" && json.Password == "admin" {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": json.User,
				"exp":      time.Now().Add(time.Hour * 1).Unix(),
			})
			tokenString, err := token.SignedString(secret)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": tokenString})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}

func SetupUserRouter(r *gin.Engine) {
	r.POST("/user/login", loginHandler)
}
