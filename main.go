package main

import (
	"project/connection"
	"project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	connection.Init()

	routes.Handler(router)
	router.Run("localhost:8080")
}
