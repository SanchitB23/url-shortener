package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"os"
	"url-shortener/database"
	"url-shortener/logger"
	"url-shortener/routes"
	"url-shortener/utils"
)

func main() {
	utils.LoadEnvs()

	logger.Init()

	database.InitDB()
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			logger.Log.Warn("Failed to close the database connection")
		}
	}(database.DB)

	server := gin.Default()

	routes.RegisterRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	logger.Log.Info("Server is running on port 10000")
	err := server.Run(":" + port)
	if err != nil {
		logger.Log.Panic("Error starting server:", err)
		return
	}
}
