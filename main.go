package main

import (
	"WhoaServer/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
