package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	server := gin.Default()

	server.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "OK"})
	})

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
