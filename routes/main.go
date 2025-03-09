package routes

import (
	"github.com/gin-gonic/gin"
	"url-shortener/config"
	"url-shortener/controllers"
	_ "url-shortener/docs" // This line is necessary for go-swagger to find your docs
	"url-shortener/middlewares"
)

func SetupRoutes(server *gin.Engine) {
	server.GET("/", controllers.RootRouteController)
	server.GET("/health", controllers.HealthCheckController)
	withAuth := server.Group("/")
	withAuth.Use(middlewares.AuthMiddleware)
	withAuth.POST("/shorten", controllers.ShortenURL)
	withAuth.GET("/:shortURL", controllers.RedirectToOriginalURL)
	config.AddSwaggerRoutes(server)
}
