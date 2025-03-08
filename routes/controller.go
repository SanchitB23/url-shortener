package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "url-shortener/docs" // This line is necessary for go-swagger to find your docs
)

// RootRouteController @Summary Redirect to health check
// @Description Redirects to the health check endpoint
// @Tags root
// @Success 301
// @Router / [get]
func RootRouteController(context *gin.Context) {
	context.Redirect(http.StatusMovedPermanently, "/health")
}

// HealthCheckController @Summary Health check
// @Description Returns the health status of the API
// @Tags health
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheckController(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "URL Shortener API is running"})
}

// @Summary Shorten URL
// @Description Shortens a given URL
// @Tags url
// @Accept json
// @Produce json
// @Param url body string true "URL to shorten"
// @Success 200 {object} map[string]string
// @Router /shorten [post]
func shortenURL(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"message": "Not Implemented"})
}

// @Summary Redirect to original URL
// @Description Redirects to the original URL based on the shortened URL
// @Tags url
// @Param shortURL path string true "Shortened URL"
// @Success 301
// @Router /{shortURL} [get]
func redirectToOriginalURL(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"message": "Not Implemented"})
}
