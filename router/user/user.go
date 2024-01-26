package user

import (
	"WhoaServer/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func loginHandler(c *gin.Context) {
	secret := []byte("secret")
	var json struct {
		User     string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var user models.User

	if c.Bind(&json) == nil {
		models.DB.Where("username = ?", json.User).First(&user)
		if models.VerifyPassword(user.Password, json.Password) {
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

func registerHandler(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = uuid.New().String()
	models.DB.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "user": newUser.ID, "username": newUser.Username})
}

func SetupUserRouter(r *gin.Engine) {
	r.POST("/user/login", loginHandler)
	r.POST("/user/register", registerHandler)
}
