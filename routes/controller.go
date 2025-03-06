package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func shortenURL(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"message": "Not Implemented"})

}

func redirectToOriginalURL(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"message": "Not Implemented"})
}
