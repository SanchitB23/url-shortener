package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"url-shortener/routes"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	err := server.Run(":" + port)
	if err != nil {
		panic("Failed to start server")
		return
	}
}
