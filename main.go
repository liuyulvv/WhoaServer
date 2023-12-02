package main

import (
	"WhoaServer/router"
	"time"

	"github.com/gin-contrib/cors"
)

func main() {
	r := router.SetupRouter()
	// r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Run(":8080")
}
