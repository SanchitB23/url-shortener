package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"os"
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/routes"
)

func SetupAndRunApp() error {
	err := config.LoadEnvs()
	if err != nil {
		return err
	}
	config.InitLogger()

	err = database.InitPostgresDB()
	if err != nil {
		return err
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			config.Log.Warn("Failed to close the database connection")
		}
	}(database.DB)

	server := gin.Default()

	routes.SetupRoutes(server)
	config.AddSwaggerRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	err = server.Run(":" + port)
	if err != nil {
		config.Log.Panic("Error starting server:", err)
		return err
	}
	config.Log.Info("Server is running on port 10000")
	return nil
}
