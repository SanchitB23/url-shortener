package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"url-shortener/logger"
	"url-shortener/routes"
	"url-shortener/utils"
)

func main() {
	utils.LoadEnvs()

	logger.Init()

	server := gin.Default()

	routes.RegisterRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	logger.Log.Info("Server is running on port 10000")
	err := server.Run(":" + port)
	if err != nil {
		logger.Log.Panic("Error starting server:", err)
		return
	}
}
