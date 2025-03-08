package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/health")
	})
	server.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "URL Shortener API is running"})
	})

	server.POST("/shorten", shortenURL)
	server.GET("/:shortURL", redirectToOriginalURL)
}
