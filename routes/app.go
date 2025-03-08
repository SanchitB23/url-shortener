package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "url-shortener/docs" // This line is necessary for go-swagger to find your docs
)

// @Summary Redirect to health check
// @Description Redirects to the health check endpoint
// @Tags root
// @Success 301
// @Router / [get]
func rootRouteController(context *gin.Context) {
	context.Redirect(http.StatusMovedPermanently, "/health")
}

// @Summary Health check
// @Description Returns the health status of the API
// @Tags health
// @Success 200 {object} map[string]string
// @Router /health [get]
func healthCheckController(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "URL Shortener API is running"})
}
