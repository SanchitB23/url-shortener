package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "url-shortener/docs" // This line is necessary for go-swagger to find your docs
	"url-shortener/logger"
	"url-shortener/models"
)

// @Summary Shorten URL
// @Description Shortens a given URL
// @Tags url
// @Accept json
// @Produce json
// @Param url body string true "URL to shorten"
// @Success 200 {object} map[string]string
// @Router /shorten [post]
func shortenURL(context *gin.Context) {

	var urlObj models.URL

	if err := context.ShouldBindJSON(&urlObj); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Log.Debug("Error binding JSON: ", err, context.JSON)
		return
	}
	logger.Log.Debugf("Host URL: %v", context.Request.Host)
	shortenedURL, err := urlObj.Shorten()
	if err != nil {
		return
	}

	// Logic to shorten the URL goes here

	context.JSON(http.StatusOK, gin.H{"shortenedURL": shortenedURL})
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
