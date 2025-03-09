package routes

import (
	"github.com/gin-gonic/gin"
	"url-shortener/controllers"
	_ "url-shortener/docs" // This line is necessary for go-swagger to find your docs
)

func SetupRoutes(server *gin.Engine) {
	server.GET("/", controllers.RootRouteController)
	server.GET("/health", controllers.HealthCheckController)
	server.POST("/shorten", controllers.ShortenURL)
	server.GET("/:shortURL", controllers.RedirectToOriginalURL)
}
